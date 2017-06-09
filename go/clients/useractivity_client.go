package clients

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

import (
	"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/session"
)

const (
	USERACTIVITY_RES_NAME = "useractivity"
)

// UserActivityClient is a client for avi UserActivity resource
type UserActivityClient struct {
	avi_session *session.AviSession
}

// NewUserActivityClient creates a new client for UserActivity resource
func NewUserActivityClient(avi_session *session.AviSession) *UserActivityClient {
	return &UserActivityClient{avi_session: avi_session}
}

func (client *UserActivityClient) GetApiPath(uuid string) string {
	path := "api/" + USERACTIVITY_RES_NAME
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of UserActivity objects
func (client *UserActivityClient) GetAll() ([]*models.UserActivity, error) {
	var plist []*models.UserActivity
	err := client.avi_session.GetCollection(client.GetApiPath(""), &plist)
	return plist, err
}

// Get an existing UserActivity by uuid
func (client *UserActivityClient) Get(uuid string) (*models.UserActivity, error) {
	var obj *models.UserActivity
	err := client.avi_session.Get(client.GetApiPath(uuid), &obj)
	return obj, err
}

// Get an existing UserActivity by name
func (client *UserActivityClient) GetByName(name string) (*models.UserActivity, error) {
	var obj *models.UserActivity
	err := client.avi_session.GetObjectByName(USERACTIVITY_RES_NAME, name, &obj)
	return obj, err
}

// Create a new UserActivity object
func (client *UserActivityClient) Create(obj *models.UserActivity) (*models.UserActivity, error) {
	var robj *models.UserActivity
	err := client.avi_session.Post(client.GetApiPath(""), obj, &robj)
	return robj, err
}

// Update an existing UserActivity object
func (client *UserActivityClient) Update(obj *models.UserActivity) (*models.UserActivity, error) {
	var robj *models.UserActivity
	path := client.GetApiPath(obj.UUID)
	err := client.avi_session.Put(path, obj, &robj)
	return robj, err
}

// Delete an existing UserActivity object with a given UUID
func (client *UserActivityClient) Delete(uuid string) error {
	return client.avi_session.Delete(client.GetApiPath(uuid))
}

// Delete an existing UserActivity object with a given name
func (client *UserActivityClient) DeleteByName(name string) error {
	res, err := client.GetByName(name)
	if err != nil {
		return err
	}
	return client.Delete(res.UUID)
}