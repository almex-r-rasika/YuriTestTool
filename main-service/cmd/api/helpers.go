package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"service/data"
	"time"
)

type SendMessageLog struct {
	MessageId  string    `json:"message_id"`
	SendTime   string    `json:"send_time"`
	SendUserId string    `json:"send_user_id"`
	Address    string    `json:"address"`
	Subject    string    `json:"subject"`
	Line1      string    `json:"line1"`
	Line2      string    `json:"line2"`
	Line3      string    `json:"line3"`
	Line4      string    `json:"line4"`
	Line5      string    `json:"line5"`
	Line6      string    `json:"line6"`
	Line7      string    `json:"line7"`
	Line8      string    `json:"line8"`
	Line9      string    `json:"line9"`
	Line10     string    `json:"line10"`
	PostTime   time.Time `json:"post_time"`
	Result     string    `json:"result"`
}

/* function for generate csv log for send messages
   @param --> SendMessageLog objects array
   @param value --> []SendMessageLog
   description --> generate csv log for send messages
   @return --> null
*/
func saveSendMessageCSV(messageLogList []SendMessageLog) error {

	// Create a new file to store CSV data
	absPath, _ := filepath.Abs("../main-service/data/log.csv")
    outputFile, err := os.Create(absPath)
    if err != nil {
		log.Fatalln(err)
        return err
    }
    defer outputFile.Close()

	// Write the header of the CSV file and the successive rows by iterating through the JSON struct array
    writer := csv.NewWriter(outputFile)
    defer writer.Flush()

    header := []string{"MessageId", "SendTime", "SendUserId", "Address", "Subject", "Line1", "Line2", "Line3", "Line4", "Line5", "Line6", "Line7", "Line8", "Line9", "Line10", "PostTime","Result"}
    if err := writer.Write(header); err != nil {
		log.Fatalln(err)
        return err
    }

	for _, l := range messageLogList {
        var csvRow []string
        csvRow = append(csvRow, l.MessageId, l.SendTime, l.SendUserId, l.Address, l.Subject, l.Line1, l.Line2, l.Line3, l.Line4, l.Line5, l.Line6, l.Line7, l.Line8, l.Line9, l.Line10, l.PostTime.GoString(), l.Result)
        if err := writer.Write(csvRow); err != nil {
			log.Fatalln(err)
            return err
        }
    }

	data.Log.Info("saved log.csv file")
    return nil
}

/* function for convert database send messages records to json objects
   @param --> null
   @param value --> null
   description --> convert send messages records into json objects
   @return --> []SendMessageLog
*/
func getSendMessageJsonList() []SendMessageLog{

	var logList []SendMessageLog
	var log SendMessageLog
	message := data.GetSendMessages()

	for _, msg := range message {
		log.MessageId = msg.MessageId
		log.SendTime = msg.SendTime
		log.SendUserId = msg.SendUserId
		log.Address = msg.Address
		log.Subject = msg.Subject
		log.Line1 = msg.Line1
		log.Line2 = msg.Line2
		log.Line3 = msg.Line3
		log.Line4 = msg.Line4
		log.Line5 = msg.Line5
		log.Line6 = msg.Line6
		log.Line7 = msg.Line7
		log.Line8 = msg.Line8
		log.Line9 = msg.Line9
		log.Line10 = msg.Line10
		log.PostTime = msg.PostTime
		log.Result = msg.Result

		logList = append(logList, log)
	}
	return logList
}