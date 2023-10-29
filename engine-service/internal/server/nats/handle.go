package rabbitmq

import (
	"encoding/json"
	"engine-service/internal/models"
	"engine-service/internal/storage/postgresql"
	"fmt"
	"log"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type handle struct {
	db  *postgresql.DB
	js  jetstream.JetStream
	jsx nats.JetStreamContext
}

func New(db *postgresql.DB, js jetstream.JetStream, jsx nats.JetStreamContext) *handle {
	return &handle{
		db:  db,
		js:  js,
		jsx: jsx,
	}
}

func (h *handle) GetEngines(cars []models.Car, connID string) {
	var (
		engines  []models.Engine
		response models.Response
	)

	response.Action = "GetCarEngines"
	response.ConnID = connID

	defer func() {
		data, err := json.Marshal(response)
		if err != nil {
			if err := h.Publish("GW.err", []byte(err.Error()), connID); err != nil {
				log.Printf("err Publish, %v", err)
			}
			return
		}

		if err := h.Publish("GW.ws", data, connID); err != nil {
			log.Printf("err Publish, %v", err)
		}
	}()

	for i := range cars {
		engine, err := h.db.GetEngine(cars[i].EngineID)
		if err != nil {
			response.Err = fmt.Errorf("err GetEngines: %w", err).Error()
			return
		}

		engines = append(engines, *engine)
	}

	response.Data = engines
}

func (h *handle) GetEngine(car models.Car, connID string) {
	var response models.Response

	response.Action = "GetCarEngine"
	response.ConnID = connID

	defer func() {
		data, err := json.Marshal(response)
		if err != nil {
			if err := h.Publish("GW.err", []byte(err.Error()), connID); err != nil {
				log.Printf("err Publish, %v", err)
			}
			return
		}

		if err := h.Publish("GW.ws", data, connID); err != nil {
			log.Printf("err Publish, %v", err)
		}
	}()

	engine, err := h.db.GetEngine(car.EngineID)
	if err != nil {
		err = fmt.Errorf("err GetEngine: %w", err)
	}

	response.Data = engine
}

func (h *handle) Worker() {
	go h.Consume()
}

func (h *handle) Consume() {
	_, err := h.jsx.Subscribe("ENGINES.*", func(m *nats.Msg) {
		err := m.Ack()
		if err != nil {
			log.Printf("Unable to Ack: %v", err)
			return
		}

		id := m.Header["conn_id"]

		var resp models.Response
		if err = json.Unmarshal(m.Data, &resp); err != nil {
			log.Printf("err Unmarshal: %v", err)
			return
		}

		switch strings.Split(m.Subject, ".")[1] {
		case "engines":
			var cars []models.Car

			s, ok := resp.Data.([]interface{})
			if !ok {
				err = fmt.Errorf("err type assertion: %w", err)
				h.Publish("GW.err", []byte(err.Error()), resp.ConnID)
				return
			}

			data, err := json.Marshal(s)
			if err != nil {
				h.Publish("GW.err", []byte(err.Error()), resp.ConnID)
				return
			}

			if err := json.Unmarshal(data, &cars); err != nil {
				h.Publish("GW.err", []byte(err.Error()), resp.ConnID)
				return
			}

			h.GetEngines(cars, id[0])

		case "engine":
			var car models.Car

			s, ok := resp.Data.(map[string]interface{})
			if !ok {
				err = fmt.Errorf("err type assertion: %w", err)
				h.Publish("GW.err", []byte(err.Error()), resp.ConnID)
				break
			}

			data, err := json.Marshal(s)
			if err != nil {
				h.Publish("GW.err", []byte(err.Error()), resp.ConnID)
				break
			}

			if err := json.Unmarshal(data, &car); err != nil {
				h.Publish("GW.err", []byte(err.Error()), resp.ConnID)
				break
			}

			h.GetEngine(car, resp.ConnID)
		}
	})

	if err != nil {
		log.Println("Subscribe failed")
		return
	}
}

func (h *handle) Publish(subj string, data []byte, connID string) error {
	_, err := h.jsx.PublishMsg(&nats.Msg{
		Data:    data,
		Subject: subj,
		Header: nats.Header{
			"conn_id": []string{connID},
		},
	})
	if err != nil {
		return fmt.Errorf("err PublishMsg: %w", err)
	}

	return nil
}
