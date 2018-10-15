package sagaQuee

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"sync"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/saga-execution-coordinator/repositories"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/saga-execution-coordinator/sagaStateMachine"
	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

type QueeMsg struct {
	State sagaStateMachine.SagaState `json:"state"`
	*repositories.BuyVideoSagaDTO
}

type SagaQuee struct {
	Ssm                  *sagaStateMachine.SagaStateMachine
	Consumer             *consumergroup.ConsumerGroup
	Topic                string
	ProducerInputChannel chan<- *sarama.ProducerMessage
	//EXAMPLE FOR KEEPING STATE
	//YOU SHOULD USE SOMETHING LIKE REDIS FOR THIS
	SuccesfulEndedSagas   sync.Map
	UnsuccesfulEndedSagas sync.Map
}

func NewSagaQuee(ssm *sagaStateMachine.SagaStateMachine, topicName string, consumer *consumergroup.ConsumerGroup, producer sarama.AsyncProducer) *SagaQuee {
	return &SagaQuee{
		Ssm:                  ssm,
		Topic:                topicName,
		Consumer:             consumer,
		ProducerInputChannel: producer.Input(),
	}
}

func (sq *SagaQuee) StartQueeProcessing() {

	for {

		var msg *sarama.ConsumerMessage
		msg = <-sq.Consumer.Messages()

		var consumedQueeMsg QueeMsg
		err := json.Unmarshal(msg.Value, &consumedQueeMsg)
		if err != nil {
			log.Println("Error Unmarshalling msg: ", err.Error())
			return
		}
		log.Println("Msg Offset: ", msg.Offset, "Msg.Timestamp: ", msg.Timestamp, "State: ", sagaStateMachine.SagaStateToString(consumedQueeMsg.State))

		if consumedQueeMsg.State == sagaStateMachine.SAGA_END || consumedQueeMsg.State == sagaStateMachine.SAGA_ROLLBACK_END || consumedQueeMsg.State == sagaStateMachine.SAGA_UNHANDLED {
			continue
		}

		nextState := sq.Ssm.ProcessSagaStateAndDecideNextState(consumedQueeMsg.State, consumedQueeMsg.BuyVideoSagaDTO)

		saramaMsg, err := sq.encodeSamaraMsg(nextState, consumedQueeMsg.BuyVideoSagaDTO)
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println(saramaMsg.Value)
		sq.ProducerInputChannel <- saramaMsg

		if nextState == sagaStateMachine.SAGA_END {
			key := sq.CreateKey(consumedQueeMsg.UserID, consumedQueeMsg.VideoID)
			sq.SuccesfulEndedSagas.Store(key, true)
		} else if nextState == sagaStateMachine.SAGA_ROLLBACK_END {
			key := sq.CreateKey(consumedQueeMsg.UserID, consumedQueeMsg.VideoID)
			sq.UnsuccesfulEndedSagas.Store(key, true)
		}

		sq.Consumer.CommitUpto(msg)
	}
}

func (sq *SagaQuee) StartSaga(bvsDTO *repositories.BuyVideoSagaDTO) error {
	saramaMsg, err := sq.encodeSamaraMsg(sagaStateMachine.SAGA_START, bvsDTO)
	if err != nil {
		return err
	}
	sq.ProducerInputChannel <- saramaMsg
	return nil
}

func (sq *SagaQuee) CreateKey(userID uint32, videoID uint32) string {
	key := strconv.Itoa(int(userID)) + "-" + strconv.Itoa(int(videoID))
	return key
}

func (sq *SagaQuee) encodeSamaraMsg(state sagaStateMachine.SagaState, bvsDTO *repositories.BuyVideoSagaDTO) (*sarama.ProducerMessage, error) {
	producedQueeMsg := &QueeMsg{
		State:           state,
		BuyVideoSagaDTO: bvsDTO,
	}

	byteMsg, err := json.Marshal(producedQueeMsg)
	if err != nil {
		return nil, errors.New("Error Marshalling msg.")
	}

	saramaMsg := &sarama.ProducerMessage{Topic: sq.Topic, Value: sarama.StringEncoder(byteMsg)}
	return saramaMsg, nil
}
