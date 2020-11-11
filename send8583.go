package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net"
    "os"
    "reflect"
    "strconv"
    "time"
    "unicode"

    "github.com/axgle/mahonia"
)

//Field represents a 8583 field
type Field struct {
    valueType int /* 0-定长 1-一位变长 2-二位变长 3-三位变长*/
    valueAttr int /* 1-右补空格 2-左补零 */
    valueLen  int
}

//Msg8583 represents a 8583 msg
type Msg8583 struct {
    Server   string
    MsgHead  string
    Field1   string
    Field2   string
    Field3   string
    Field4   string
    Field5   string
    Field6   string
    Field7   string
    Field8   string
    Field9   string
    Field10  string
    Field11  string
    Field12  string
    Field13  string
    Field14  string
    Field15  string
    Field16  string
    Field17  string
    Field18  string
    Field19  string
    Field20  string
    Field21  string
    Field22  string
    Field23  string
    Field24  string
    Field25  string
    Field26  string
    Field27  string
    Field28  string
    Field29  string
    Field30  string
    Field31  string
    Field32  string
    Field33  string
    Field34  string
    Field35  string
    Field36  string
    Field37  string
    Field38  string
    Field39  string
    Field40  string
    Field41  string
    Field42  string
    Field43  string
    Field44  string
    Field45  string
    Field46  string
    Field47  string
    Field48  string
    Field49  string
    Field50  string
    Field51  string
    Field52  string
    Field53  string
    Field54  string
    Field55  string
    Field56  string
    Field57  string
    Field58  string
    Field59  string
    Field60  string
    Field61  string
    Field62  string
    Field63  string
    Field64  string
    Field65  string
    Field66  string
    Field67  string
    Field68  string
    Field69  string
    Field70  string
    Field71  string
    Field72  string
    Field73  string
    Field74  string
    Field75  string
    Field76  string
    Field77  string
    Field78  string
    Field79  string
    Field80  string
    Field81  string
    Field82  string
    Field83  string
    Field84  string
    Field85  string
    Field86  string
    Field87  string
    Field88  string
    Field89  string
    Field90  string
    Field91  string
    Field92  string
    Field93  string
    Field94  string
    Field95  string
    Field96  string
    Field97  string
    Field98  string
    Field99  string
    Field100 string
    Field101 string
    Field102 string
    Field103 string
    Field104 string
    Field105 string
    Field106 string
    Field107 string
    Field108 string
    Field109 string
    Field110 string
    Field111 string
    Field112 string
    Field113 string
    Field114 string
    Field115 string
    Field116 string
    Field117 string
    Field118 string
    Field119 string
    Field120 string
    Field121 string
    Field122 string
    Field123 string
    Field124 string
    Field125 string
    Field126 string
    Field127 string
    Field128 string
}

var fieldDesc = []Field{
    Field{0, 2, 5}, Field{0, 1, 5}, Field{0, 1, 5}, Field{0, 2, 9},
    Field{0, 2, 8}, Field{0, 2, 6}, Field{0, 1, 6}, Field{0, 1, 6},
    Field{3, 1, 27}, Field{0, 1, 12}, Field{0, 1, 1}, Field{0, 1, 4},
    Field{0, 1, 50}, Field{0, 2, 8}, Field{3, 1, 127}, Field{0, 1, 4},
    Field{0, 2, 6}, Field{0, 1, 4}, Field{3, 1, 17}, Field{0, 1, 1},
    Field{0, 1, 2}, Field{0, 1, 2}, Field{0, 1, 3}, Field{0, 1, 3},
    Field{0, 1, 60}, Field{0, 1, 60}, Field{0, 1, 60}, Field{0, 2, 8},
    Field{0, 2, 8}, Field{0, 1, 24}, Field{0, 1, 24}, Field{0, 1, 24},
    Field{0, 1, 24}, Field{0, 2, 4}, Field{0, 2, 4}, Field{0, 2, 4},
    Field{0, 1, 24}, Field{0, 1, 24}, Field{0, 2, 16}, Field{0, 2, 16},
    Field{0, 2, 16}, Field{0, 2, 16}, Field{0, 2, 16}, Field{0, 2, 8},
    Field{0, 2, 8}, Field{0, 2, 8}, Field{0, 2, 8}, Field{0, 2, 4},
    Field{0, 2, 4}, Field{0, 2, 3}, Field{0, 2, 3}, Field{0, 2, 12},
    Field{0, 2, 10}, Field{0, 2, 10}, Field{0, 2, 10}, Field{0, 1, 19},
    Field{0, 1, 19}, Field{0, 1, 16}, Field{0, 1, 16}, Field{0, 1, 16},
    Field{0, 1, 16}, Field{0, 1, 20}, Field{0, 1, 20}, Field{0, 1, 10},
    Field{0, 1, 10}, Field{0, 1, 1}, Field{0, 1, 1}, Field{0, 1, 1},
    Field{0, 1, 1}, Field{0, 1, 1}, Field{0, 1, 1}, Field{0, 1, 1},
    Field{0, 1, 10}, Field{0, 1, 76}, Field{0, 1, 37}, Field{0, 1, 141},
    Field{0, 2, 9}, Field{0, 2, 9}, Field{0, 1, 16}, Field{0, 1, 16},
    Field{0, 1, 60}, Field{0, 1, 60}, Field{0, 1, 60}, Field{0, 1, 8},
    Field{0, 1, 8}, Field{0, 1, 20}, Field{0, 1, 20}, Field{0, 1, 20},
    Field{0, 1, 3}, Field{0, 1, 3}, Field{0, 1, 5}, Field{0, 1, 6},
    Field{0, 1, 1}, Field{0, 1, 8}, Field{0, 1, 600}, Field{3, 1, 691},
    Field{0, 1, 8}, Field{0, 1, 1}, Field{0, 1, 1}, Field{3, 1, 112},
    Field{3, 1, 205}, Field{3, 1, 412}, Field{3, 1, 173}, Field{3, 1, 0},
    Field{3, 1, 174}, Field{3, 1, 24}, Field{3, 1, 279}, Field{3, 1, 359},
    Field{3, 1, 270}, Field{3, 1, 0}, Field{3, 1, 0}, Field{3, 1, 154},
    Field{3, 1, 63}, Field{3, 1, 59}, Field{2, 1, 59}, Field{3, 1, 162},
    Field{3, 1, 17}, Field{3, 1, 90}, Field{3, 1, 90}, Field{3, 1, 162},
    Field{3, 1, 162}, Field{3, 1, 260}, Field{3, 1, 999}, Field{3, 1, 226},
    Field{3, 1, 281}, Field{3, 1, 43}, Field{0, 1, 1}, Field{0, 1, 8},
}

func main() {
    var m8583 Msg8583
    var msg8583 []Msg8583

    var isFind bool
    var strTranCode string

    // jsonTranInfo, err := ioutil.ReadFile("msg.json")
    // CheckErr(err)

    // CheckErr(json.Unmarshal(jsonTranInfo, &msg8583))
    // m8583 = msg8583[0]
    // sendMsg := EncodeMsg(m8583)
    // fmt.Println("Send msg ---> length =", len(sendMsg))
    // fmt.Println(sendMsg)
    // recvMsg := SendMsg(sendMsg, m8583.Server)
    // fmt.Println("Recv msg ---> length =", len(recvMsg))
    // fmt.Println(recvMsg)
    // DecodeMsg(recvMsg[76:])
    // fmt.Println("=======================================================")

    for {
        jsonTranInfo, err := ioutil.ReadFile("msg.json")
        CheckErr(err)

        err = json.Unmarshal(jsonTranInfo, &msg8583)
        if err != nil {
            fmt.Println("\nmsg.json err")
            fmt.Print("\n(r)-Reload (q)-Quit : ")
            _, err = fmt.Scanln(&strTranCode)
            if err != nil || len(strTranCode) <= 0 {
                fmt.Println("Input error!!!")
                fmt.Println("=======================================================")
                continue
            }

            if strTranCode == "r" {
                fmt.Println("=======================================================")
                continue
            }

            if strTranCode == "q" {
                break
            }

        }
        /* 不打印
        ** for i := 0; i < len(msg8583); i++ {
        ** fmt.Printf("%02d: [%s]\n", i, msg8583[i].Field16)
        ** }
        */

        fmt.Print("Welcome to use 8583 message tool.\n(r)-Reload (q)-Quit\nTranCode:")
        _, err = fmt.Scanln(&strTranCode)
        if err != nil || len(strTranCode) <= 0 {
            fmt.Println("Input error!!!")
            fmt.Println("=======================================================")
            continue
        }

        if strTranCode == "r" {
            fmt.Println("=======================================================")
            continue
        }

        if strTranCode == "q" {
            break
        }

        isFind = false
        for i := 0; i < len(msg8583); i++ {
            if msg8583[i].Field16 == strTranCode {
                m8583 = msg8583[i]
                isFind = true
            }
        }

        if isFind == false {
            fmt.Println("Input error!!!")
            fmt.Println("=======================================================")
            continue
        }

        sendMsg := EncodeMsg(m8583)
        fmt.Println("Send msg ---> length =", len(sendMsg))
        fmt.Println(sendMsg)
        recvMsg, err := SendMsg(sendMsg, m8583.Server)
        if err != nil {
            fmt.Println(err.Error())
            fmt.Println("=======================================================")
            continue
        }
        fmt.Println("Recv msg ---> length =", len(recvMsg))
        fmt.Println(recvMsg)
        DecodeMsg(recvMsg[76:])
        fmt.Println("=======================================================")
    }
}

//EncodeMsg is to encode the 8583 msg
func EncodeMsg(msg8583 Msg8583) string {
    var strTmp string
    var strMsg string
    var bBitMap [128]byte

    //fmt.Printf("Server: [%s] len=[%d]\n", msg8583.Server, len(msg8583.Server))
    //fmt.Printf("MsgHead: [%s] len=[%d]\n", msg8583.MsgHead, len(msg8583.MsgHead))

    InitArray(bBitMap[:])

    bBitMap[0] = '1'
    for i := 1; i < 128; i++ {
        strFieldName := fmt.Sprintf("Field%d", i+1)
        //fmt.Println(strFieldName)
        strValue := mahonia.NewEncoder("gbk").ConvertString(GetValueByName(&msg8583, strFieldName))
        //fmt.Printf("strValue=[%s]\n", strValue)
        if len(strValue) != 0 {
            bBitMap[i] = '1'

            if fieldDesc[i].valueType == 0 { /* 0-定长 1-一位变长 2-二位变长 3-三位变长 */
                if fieldDesc[i].valueAttr == 1 { /* 1-右补空格 2-左补零 */
                    strTmp += string(fmt.Sprintf("%-*s", fieldDesc[i].valueLen, strValue)[:fieldDesc[i].valueLen])
                } else if fieldDesc[i].valueAttr == 2 { /* 1-右补空格 2-左补零 */
                    //strTmp += fmt.Sprintf("%0*s", fieldDesc[i].valueLen, strValue)
                    strTmp += string(fmt.Sprintf("%0*s", fieldDesc[i].valueLen, strValue)[:fieldDesc[i].valueLen])
                }
            } else if fieldDesc[i].valueType == 1 { /* 0-定长 1-一位变长 2-二位变长 3-三位变长 */
                strTmp += fmt.Sprintf("%01d", fieldDesc[i].valueLen)
                //strTmp += fmt.Sprintf("%s", strValue)
                //strTmp += fmt.Sprintf("%-*s", fieldDesc[i].valueLen, strValue)
                strTmp += string(fmt.Sprintf("%-*s", fieldDesc[i].valueLen, strValue)[:fieldDesc[i].valueLen])
            } else if fieldDesc[i].valueType == 2 { /* 0-定长 1-一位变长 2-二位变长 3-三位变长 */
                strTmp += fmt.Sprintf("%02d", fieldDesc[i].valueLen)
                //strTmp += fmt.Sprintf("%s", strValue)
                //strTmp += fmt.Sprintf("%-*s", fieldDesc[i].valueLen, strValue)
                strTmp += string(fmt.Sprintf("%-*s", fieldDesc[i].valueLen, strValue)[:fieldDesc[i].valueLen])
            } else if fieldDesc[i].valueType == 3 { /* 0-定长 1-一位变长 2-二位变长 3-三位变长 */
                strTmp += fmt.Sprintf("%03d", fieldDesc[i].valueLen)
                //strTmp += fmt.Sprintf("%s", strValue)
                //strTmp += fmt.Sprintf("%-*s", fieldDesc[i].valueLen, strValue)
                strTmp += string(fmt.Sprintf("%-*s", fieldDesc[i].valueLen, strValue)[:fieldDesc[i].valueLen])
            }
        }
    }
    // for i := 1; i < 129; i++ {
    //     fmt.Printf("bBitMap[%d]=[%c]\n", i, bBitMap[i-1])
    // }

    strBitMap := GetBitMapString(bBitMap)
    strMsg += msg8583.MsgHead
    strMsg += strBitMap
    strMsg += strTmp

    strMsgLen := fmt.Sprintf("%08d", len(strMsg))

    return strMsgLen + strMsg
}

//DecodeMsg is to encode the 8583 msg
func DecodeMsg(msg string) {
    var offset int

    strBitMap, lenBitMap := GetBitMapBinary(msg[0:16])
    offset += lenBitMap

    for i := 1; i < lenBitMap*8; i++ {
        if strBitMap[i] == '1' {
            if fieldDesc[i].valueType == 0 { /* 0-定长 1-一位变长 2-二位变长 3-三位变长*/
                fmt.Printf("Field%d:[%s]\n", i+1, mahonia.NewDecoder("gbk").ConvertString(string(msg[offset:offset+fieldDesc[i].valueLen])))
                offset += fieldDesc[i].valueLen
            } else if fieldDesc[i].valueType == 1 { /* 0-定长 1-一位变长 2-二位变长 3-三位变长*/
                tmpLen, _ := strconv.Atoi(msg[offset : offset+1])
                offset++
                fmt.Printf("Field%d:[%s]\n", i+1, mahonia.NewDecoder("gbk").ConvertString(string(msg[offset:offset+tmpLen])))
                offset += tmpLen
            } else if fieldDesc[i].valueType == 2 { /* 0-定长 1-一位变长 2-二位变长 3-三位变长*/
                tmpLen, _ := strconv.Atoi(msg[offset : offset+2])
                offset += 2
                fmt.Printf("Field%d:[%s]\n", i+1, mahonia.NewDecoder("gbk").ConvertString(string(msg[offset:offset+tmpLen])))
                offset += tmpLen
            } else if fieldDesc[i].valueType == 3 { /* 0-定长 1-一位变长 2-二位变长 3-三位变长*/
                tmpLen, _ := strconv.Atoi(msg[offset : offset+3])
                offset += 3
                fmt.Printf("Field%d:[%s]\n", i+1, mahonia.NewDecoder("gbk").ConvertString(string(msg[offset:offset+tmpLen])))
                offset += tmpLen
            }

        }
    }
}

//GetValueByName is used to get the value by the field name sting
func GetValueByName(m interface{}, f string) string {
    return reflect.Indirect(reflect.ValueOf(m)).FieldByName(f).String()
}

//InitArray is used to init the array
func InitArray(a []byte) {
    for i := 0; i < len(a); i++ {
        a[i] = '0'
    }
}

//GetBitMapString is used to get the BitMap string
func GetBitMapString(a [128]byte) string {
    var tmpArray [16]byte
    for i := 0; i < 16; i++ {
        tmpValue, err := strconv.ParseUint(string(a[i*8:(i+1)*8]), 2, 8)
        CheckErr(err)
        tmpArray[i] = byte(tmpValue)
    }

    return string(tmpArray[:])
}

//GetBitMapBinary is used to get the BitMap string
func GetBitMapBinary(s string) (string, int) {
    var tmpString string
    for i := 0; i < 16; i++ {
        tmpString += fmt.Sprintf("%08s", strconv.FormatUint(uint64(s[i]), 2))
    }

    if tmpString[0] == '0' {
        return tmpString[:64], 8
    }

    return tmpString, 16
}

//SendMsg is used to send the msg to the svr
func SendMsg(reqMsg string, svr string) (string, error) {
    conn, err := net.DialTimeout("tcp", svr, 5*time.Second)
    if err != nil {
        return "", err
    }
    defer conn.Close()

    conn.SetDeadline(time.Now().Add(5 * time.Second))

    conn.Write([]byte(reqMsg))

    var rspMsg = make([]byte, 1024*1024)

    nRead, err := conn.Read(rspMsg)
    if err != nil {
        return "", err
    }

    return string(rspMsg[0:nRead]), nil
}

func IsHaveChinese(str string) bool {
    for _, v := range str {
        if unicode.Is(unicode.Han, v) {
            return true
        }
    }
    return false
}

//CheckErr is used to check the err
func CheckErr(err error) {
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(-1)
    }
}
