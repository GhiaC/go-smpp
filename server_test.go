package go_smpp

import (
	"testing"
)

func TestReadMore(t *testing.T) {
	data1 := []byte{0x00, 0x00, 0x00, 0x2F, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00}
	data2 := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x53, 0x4D, 0x50, 0x50, 0x33, 0x54, 0x45, 0x53, 0x54, 0x00,
		0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x30, 0x38, 0x00, 0x53, 0x55, 0x42, 0x4D, 0x49, 0x54, 0x31, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x2F, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00}
	data3 := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x53, 0x4D, 0x50, 0x50, 0x33, 0x54, 0x45, 0x53, 0x54, 0x00,
		0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x30, 0x38, 0x00, 0x53, 0x55, 0x42, 0x4D, 0x49, 0x54, 0x31, 0x00, 0x00, 0x01, 0x01, 0x00}

	buf := make([]byte, 20480)
	read := make([]byte, 10240)
	more := make([]byte, 20480)
	more = more[:0]
	buf = buf[:0]

	var smppPdu Pdu

	read = data1
	smppPdu = readMore(&buf, &more, &read, len(read))

	if smppPdu.complete == true {
		t.Error("Expected incomplete PDU got", smppPdu.complete)
	}

	read = data2
	smppPdu = readMore(&buf, &more, &read, len(read))

	if smppPdu.commandLength != 47 {
		t.Error("Expected pdu.command_length 47", smppPdu.commandLength)
	}
	if smppPdu.complete != true {
		t.Error("Expected complete PDU got", smppPdu.complete)
	}

	read = data3
	smppPdu = readMore(&buf, &more, &read, len(read))

	if smppPdu.commandLength != 47 {
		t.Error("Expected pdu.command_length 47", smppPdu.commandLength)
	}
	if smppPdu.complete != true {
		t.Error("Expected complete PDU got", smppPdu.complete)
	}

}
