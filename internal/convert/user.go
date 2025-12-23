package convert

import (
	"github.com/JeyKeyAlex/TourProject/internal/entities"

	pb "github.com/JeyKeyAlex/TourProject-proto/go-genproto/user"
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
