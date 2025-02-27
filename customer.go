package zammad

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type CustomerID struct {
	Value interface{} `json:"-"`
}

func (c *CustomerID) String() string {
	if c.Value == nil {
		return ""
	}
	return fmt.Sprintf("%v", c.Value)
}

func (c CustomerID) Int() int {
	if c.Value == nil {
		return -1
	}

	i, ok := c.Value.(int)
	if ok {
		return i
	}

	s, ok := c.Value.(string)
	if !ok {
		return -1
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}

	return i
}

func (c *CustomerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Value)
}

func (c *CustomerID) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value == nil {
		return nil
	}

	// JSON numbers are decoded as float64
	if i, ok := value.(float64); ok {
		c.Value = int(i)
		return nil
	}

	c.Value = value
	return nil
}
