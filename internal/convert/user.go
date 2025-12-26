package convert

import (
	"github.com/JeyKeyAlex/TourProject/internal/entities"

	pb "github.com/JeyKeyAlex/TestProject-genproto/user"
)

func GetUserEntityToEntry(eResp *entities.User) (*pb.User, error) {
	user := &pb.User{}

	user.Id = eResp.Id
	user.Name = eResp.Name
	user.LastName = eResp.LastName
	user.MiddleName = eResp.MiddleName
	user.Nickname = eResp.Nickname
	user.Email = eResp.Email
	user.PhoneNumber = eResp.PhoneNumber

	createdAt := eResp.CreatedAt.String()
	user.CreatedAt = &createdAt

	return user, nil
}

func UpdateUserEntryToEntity(request *pb.UpdateUserRequest) (*entities.UpdateUserRequest, error) {
	user := &entities.UpdateUserRequest{}

	user.Id = request.Id
	user.Name = request.Name
	user.LastName = request.LastName
	user.MiddleName = request.MiddleName
	user.Nickname = request.Nickname
	user.Email = request.Email
	user.PhoneNumber = request.PhoneNumber

	return user, nil
}
