package canvaslms

import (
	"encoding/json"
	"strconv"
)

// GetCourses - Returns a slice of course
func (c *Client) GetCourses(courseCodes []int) ([]Course, error) {
	var courses []Course

	for _, courseCode := range courseCodes {
		course, err := c.GetCourse(courseCode)
		if err != nil {
			return nil, err
		}
		courses = append(courses, *course)

	}
	return courses, nil
}

// GetCourse - Returns a course
func (c *Client) GetCourse(courseCode int) (*Course, error) {
	resp, err := c.RESTClient.R().Get("/api/v1/courses/" + strconv.Itoa(courseCode))
	if err != nil {
		return nil, err
	}

	resultsJson := resp.Body()

	course := Course{}
	err = json.Unmarshal(resultsJson, &course)
	if err != nil {
		return nil, err
	}

	course.Teachers, err = getUsers(c, course.Id, "teacher")
	if err != nil {
		return nil, err
	}
	course.TAs, err = getUsers(c, course.Id, "ta")
	if err != nil {
		return nil, err
	}
	course.Students, err = getUsers(c, course.Id, "student")
	if err != nil {
		return nil, err
	}

	course.Groups, err = getGroups(c, course.Id)
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func getUsers(client *Client, courseCode int, enrollmentType string) ([]User, error) {
	client.RESTClient.
		SetQueryParams(map[string]string{
			"enrollment_type[]": enrollmentType,
		})

	resultsJson, err := client.getResults("/api/v1/courses/" + strconv.Itoa(courseCode) + "/users")
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(resultsJson, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getGroups(client *Client, courseCode int) ([]Group, error) {

	resultsJson, err := client.getResults("/api/v1/courses/" + strconv.Itoa(courseCode) + "/groups")
	if err != nil {
		return nil, err
	}

	var groups []Group
	err = json.Unmarshal(resultsJson, &groups)
	if err != nil {
		return nil, err
	}

	for i, group := range groups {
		usersJson, err := client.getResults("/api/v1/groups/" + strconv.Itoa(group.Id) + "/users")
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(usersJson, &groups[i].Students)
		if err != nil {
			return nil, err
		}
	}

	return groups, nil
}
