package database

import (
	"github.com/divyag9/gomodel/pkg/pb"
	ora "gopkg.in/rana/ora.v4"
)

//Info contains information required for the database
type Info struct {
	Session *ora.Ses
}

//GetImageDetailsByOrderNumber retrieves the imageDetails for a given orderNumber
func (i *Info) GetImageDetailsByOrderNumber(orderNumber int64) ([]*contentservice.ImageDetail, error) {
	return []*contentservice.ImageDetail{&contentservice.ImageDetail{}}, nil
}

//GetImageDetailsByImageIds retrieves the imageDetails for a given set og imageIds
func (i *Info) GetImageDetailsByImageIds(imageIds []int64) ([]*contentservice.ImageDetail, error) {
	return []*contentservice.ImageDetail{&contentservice.ImageDetail{}}, nil
}
