package main

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	_ "github.com/emersion/go-message/charset"
	"github.com/emersion/go-message/mail"
	"github.com/k3a/html2text"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS("imap.qq.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login("844870676@qq.com", "密码"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	//// List mailboxes
	//mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	//go func() {
	//	done <- c.List("", "*", mailboxes)
	//}()
	//
	//log.Println("Mailboxes:")
	//for m := range mailboxes {
	//	log.Println("* " + m.Name)
	//}
	//
	//if err := <-done; err != nil {
	//	log.Fatal(err)
	//}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only subtract if the result is > 0
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	section := imap.BodySectionName{}

	go func() {
		// imap.FetchEnvelope,imap.FetchBody,imap.FetchBodyStructure,imap.FetchRFC822Text
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}, messages)
	}()

	log.Println("Last 4 messages:")

	receMsgs := map[uint32]*imap.Message{}

	for msg := range messages {
		receMsgs[msg.SeqNum] = msg

		if msg.Envelope.Subject != "Fwd: VIP feedback" {
			//if msg.Envelope.Subject != "测试邮件解析-正文图片" {
			continue
		}
		log.Println("* " + msg.Envelope.Subject)

		r := msg.GetBody(&section)
		if r == nil {
			log.Println(msg)
			continue
		}
		m, err := mail.CreateReader(r)
		if err != nil {
			continue
		}

		hf := m.Header.Fields()
		hs := ""
		for hf.Next() {
			t, err := hf.Text()
			if err != nil {
				log.Println(nil, "fetch 403:", err)
				continue
			}
			hs += t + "\r\n"
		}

		var textBody []byte
		var htmlBody []byte
		attachments := []at{}

		// Process each message's part
		for {
			p, err := m.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Println(nil, "fetch 423:", err)
				break
			}

			switch h := p.Header.(type) {
			case *mail.InlineHeader:
				// This is the message's text (can be plain-text or HTML)
				b, err := ioutil.ReadAll(p.Body)
				if err != nil {
					continue
				}

				ct, _, err := h.Header.ContentType()
				if err != nil {
					continue
				}
				//filename, err := h.Filename()
				if strings.Contains(ct, "image") {
					//log.Println("============>" ,ct)
					contentType := p.Header.Get("Content-Type") // image/jpeg; name="image1.jpeg"
					index := strings.Index(contentType, "\"")
					fileName := contentType[index+1 : len(contentType)-1]
					log.Println("---------------FileName:", fileName)

					attachments = append(attachments, at{
						filename:    fileName,
						contentType: ct,
						body:        b,
					})
				}
				if ct == "text/plain" {
					textBody = b
				} else if ct == "text/html" {
					htmlBody = b
				}
			case *mail.AttachmentHeader:
				// This is an attachment
				filename, err := h.Filename()
				log.Println("=======================>fileName=", filename)
				if err != nil {
					continue
				}
				ct, _, err := h.Header.ContentType()
				if err != nil {
					continue
				}
				b, err := ioutil.ReadAll(p.Body)
				if err != nil {
					continue
				}
				attachments = append(attachments, at{
					filename:    filename,
					contentType: ct,
					body:        b,
				})
			}
		}

		if textBody == nil && htmlBody != nil {
			textBody = []byte(html2text.HTML2Text(string(htmlBody)))
		} else if textBody != nil && htmlBody == nil {
			htmlBody = textBody
		} else if textBody == nil && htmlBody == nil {
			textBody = []byte{}
			htmlBody = []byte{}
		}

		log.Println(len(receMsgs))
		log.Println(string(htmlBody))
		log.Println(string(textBody))
		log.Println("attach list size = ", len(attachments))

	} //end

	if err := <-done; err != nil {
		log.Fatal(err)
	}
	log.Println("Done!")
}

type at struct {
	filename    string
	contentType string
	body        []byte
}
