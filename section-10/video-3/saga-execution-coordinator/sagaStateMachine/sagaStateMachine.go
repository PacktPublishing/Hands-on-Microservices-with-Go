package sagaStateMachine

import (
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/saga-execution-coordinator/repositories"
)

type SagaState uint16

const (
	//The usage of underscores is not idiomatic
	//But I think for this example it will make things clearer
	SAGA_START SagaState = iota
	INSERT_BOUGHT_VIDEO_START
	INSERT_BOUGHT_VIDEO_END
	UPDATE_USER_ACCOUNT_START
	UPDATE_USER_ACCOUNT_END
	UPDATE_AGENT_ACCOUNT_START
	UPDATE_AGENT_ACCOUNT_END
	SAGA_END

	//From End to Begining because it's clearer
	UPDATE_AGENT_ACCOUNT_ROLLBACK_END
	UPDATE_USER_ACCOUNT_ROLLBACK_START
	UPDATE_USER_ACCOUNT_ROLLBACK_END
	INSERT_BOUGHT_VIDEO_ROLLBACK_START
	INSERT_BOUGHT_VIDEO_ROLLBACK_END
	SAGA_ROLLBACK_END

	UPDATE_AGENT_ACCOUNT_ROLLBACK_FAILED
	UPDATE_USER_ACCOUNT_ROLLBACK_FAILED
	INSERT_BOUGHT_VIDEO_ROLLBACK_FAILED

	SAGA_UNKNOWN_STATE
	SAGA_UNHANDLED
)

type SagaStateMachine struct {
	VideosRepo *repositories.RestVideosRepository
	UsersRepo  *repositories.RestUsersRepository
	AgentsRepo *repositories.RestAgentsRepository
}

func (ssm *SagaStateMachine) ProcessSagaStateAndDecideNextState(currentState SagaState, bvsDTO *repositories.BuyVideoSagaDTO) SagaState {

	switch currentState {
	case SAGA_START:
		return INSERT_BOUGHT_VIDEO_START

	case INSERT_BOUGHT_VIDEO_START:
		err := ssm.InsertBoughtVideo(bvsDTO)
		if err != nil {
			log.Println(err.Error())
			return INSERT_BOUGHT_VIDEO_ROLLBACK_END
		}
		return INSERT_BOUGHT_VIDEO_END

	case INSERT_BOUGHT_VIDEO_END:
		return UPDATE_USER_ACCOUNT_START

	case UPDATE_USER_ACCOUNT_START:
		err := ssm.UpateUserAccount(bvsDTO)
		if err != nil {
			log.Println(err.Error())
			return UPDATE_USER_ACCOUNT_ROLLBACK_END
		}
		return UPDATE_USER_ACCOUNT_END

	case UPDATE_USER_ACCOUNT_END:
		return UPDATE_AGENT_ACCOUNT_START

	case UPDATE_AGENT_ACCOUNT_START:
		err := ssm.UpateAgentAccount(bvsDTO)
		if err != nil {
			log.Println(err.Error())
			return UPDATE_AGENT_ACCOUNT_ROLLBACK_END
		}
		return UPDATE_AGENT_ACCOUNT_END

	case UPDATE_AGENT_ACCOUNT_END:
		return SAGA_END

	case UPDATE_AGENT_ACCOUNT_ROLLBACK_END:
		return UPDATE_USER_ACCOUNT_ROLLBACK_START

	case UPDATE_USER_ACCOUNT_ROLLBACK_START:
		err := ssm.UpateUserAccountRollback(bvsDTO)
		if err != nil {
			log.Println(err.Error())
			return UPDATE_USER_ACCOUNT_ROLLBACK_FAILED
		}
		return UPDATE_USER_ACCOUNT_ROLLBACK_END

	case UPDATE_USER_ACCOUNT_ROLLBACK_END:
		return INSERT_BOUGHT_VIDEO_ROLLBACK_START

	case INSERT_BOUGHT_VIDEO_ROLLBACK_START:
		err := ssm.InsertBoughtVideoRollback(bvsDTO)
		if err != nil {
			log.Println(err.Error())
			return INSERT_BOUGHT_VIDEO_ROLLBACK_FAILED
		}
		return INSERT_BOUGHT_VIDEO_ROLLBACK_END

	case INSERT_BOUGHT_VIDEO_ROLLBACK_END:
		return SAGA_ROLLBACK_END

	case UPDATE_USER_ACCOUNT_ROLLBACK_FAILED:
	case INSERT_BOUGHT_VIDEO_ROLLBACK_FAILED:
		return SAGA_UNHANDLED
		//Harsh error condition.
		//Not implemented.
		//You could try retrying after some time for example.
		//Or logging it to another quee for processing once that
		//part of your system is up again.

	case SAGA_UNKNOWN_STATE:
		return SAGA_UNHANDLED
	}

	//UNKNOWN STATE
	return SAGA_UNKNOWN_STATE
}

func (ssm *SagaStateMachine) InsertBoughtVideo(bvsDTO *repositories.BuyVideoSagaDTO) error {
	return ssm.VideosRepo.InsertBoughtVideo(bvsDTO)
}

func (ssm *SagaStateMachine) UpateUserAccount(bvsDTO *repositories.BuyVideoSagaDTO) error {
	return ssm.UsersRepo.UpdateUserAccount(bvsDTO)
}

func (ssm *SagaStateMachine) UpateAgentAccount(bvsDTO *repositories.BuyVideoSagaDTO) error {
	return ssm.AgentsRepo.UpdateAgentAccount(bvsDTO)
}

func (ssm *SagaStateMachine) UpateAgentAccountRollback(bvsDTO *repositories.BuyVideoSagaDTO) error {
	return ssm.AgentsRepo.RollbackUpdateAgentAccount(bvsDTO)
}

func (ssm *SagaStateMachine) UpateUserAccountRollback(bvsDTO *repositories.BuyVideoSagaDTO) error {
	return ssm.UsersRepo.RollbackUpdateUserAccount(bvsDTO)
}

func (ssm *SagaStateMachine) InsertBoughtVideoRollback(bvsDTO *repositories.BuyVideoSagaDTO) error {
	return ssm.VideosRepo.DeleteBoughtVideo(bvsDTO)
}

func SagaStateToString(state SagaState) string {
	switch state {
	case SAGA_START:
		return "SAGA_START"
	case INSERT_BOUGHT_VIDEO_START:
		return "INSERT_BOUGHT_VIDEO_START"
	case INSERT_BOUGHT_VIDEO_END:
		return "INSERT_BOUGHT_VIDEO_END"
	case UPDATE_USER_ACCOUNT_START:
		return "UPDATE_USER_ACCOUNT_START"
	case UPDATE_USER_ACCOUNT_END:
		return "UPDATE_USER_ACCOUNT_END"
	case UPDATE_AGENT_ACCOUNT_START:
		return "UPDATE_AGENT_ACCOUNT_START"
	case UPDATE_AGENT_ACCOUNT_END:
		return "UPDATE_AGENT_ACCOUNT_END"
	case SAGA_END:
		return "SAGA_END"

	case SAGA_UNKNOWN_STATE:
		return "SAGA_UNKNOWN_STATE"

	case UPDATE_AGENT_ACCOUNT_ROLLBACK_END:
		return "UPDATE_AGENT_ACCOUNT_ROLLBACK_END"
	case UPDATE_USER_ACCOUNT_ROLLBACK_START:
		return "UPDATE_USER_ACCOUNT_ROLLBACK_START"
	case UPDATE_USER_ACCOUNT_ROLLBACK_END:
		return "UPDATE_USER_ACCOUNT_ROLLBACK_END"
	case INSERT_BOUGHT_VIDEO_ROLLBACK_START:
		return "INSERT_BOUGHT_VIDEO_ROLLBACK_START"
	case INSERT_BOUGHT_VIDEO_ROLLBACK_END:
		return "INSERT_BOUGHT_VIDEO_ROLLBACK_END"
	case SAGA_ROLLBACK_END:
		return "SAGA_ROLLBACK_END"

	case UPDATE_AGENT_ACCOUNT_ROLLBACK_FAILED:
		return "UPDATE_AGENT_ACCOUNT_ROLLBACK_FAILED"
	case UPDATE_USER_ACCOUNT_ROLLBACK_FAILED:
		return "UPDATE_USER_ACCOUNT_ROLLBACK_FAILED"
	case INSERT_BOUGHT_VIDEO_ROLLBACK_FAILED:
		return "INSERT_BOUGHT_VIDEO_ROLLBACK_FAILED"
	case SAGA_UNHANDLED:
		return "SAGA_UNHANDLED"
	}

	//UNKNOWN STATE
	return "SAGA_UNKNOWN_STATE"
}
