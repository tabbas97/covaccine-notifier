package main

import (
	"net/smtp"
	"strconv"
)

func sendMail(id, pass string, age int, body string) error {
	msg := "From: " + id + "\n" +
		"To: " + id + "\n" +
		"Subject: AGE: " + strconv.Itoa(age) + " Vaccination slots are available\n\n" +
		"Vaccination slots are available at the following centers:\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", id, pass, "smtp.gmail.com"),
		id, []string{id}, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}
