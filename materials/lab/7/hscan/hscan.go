package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"sync"
	"os"
)

//==========================================================================\\

var shalookup =  make(map[string]string)
//shalookup = make(map[string]string)
var md5lookup = make(map[string]string)
//md5lookup := make(map[string]string
var mutexmd5 = &sync.Mutex{}
var mutexsha = &sync.Mutex{} 
var wg sync.WaitGroup

func GuessSingle(sourceHash string, filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		// TODO - From the length of the hash you should know which one of these to check ...
		// add a check and logicial structure
		//show := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		//fmt.Printf("Hash value: %s", show)
		length := len(sourceHash)
		if length == 32 {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
			}
		}else if length == 64 {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func GenHashMaps(filename string) {

	//TODO
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time
	//OPTIONAL -- Can you use workers to make this even faster


	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password := scanner.Text()
		wg.Add(2)
		
		go genMD5Hash(password)
		go genSHAHash(password)
	}
	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)
}

func genMD5Hash(password string) {  
	defer wg.Done()
    hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	
	// md5lookup = map[string]string{           // Map literal
	// 	"key":  hash,
	// 	"value": password,
	// }
	mutexmd5.Lock()
	md5lookup[hash] = password
	mutexmd5.Unlock()
	
	
}

func genSHAHash(password string) {  
	defer wg.Done()
    hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	mutexsha.Lock()
	shalookup[hash] = password
	mutexsha.Unlock()
}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	//fmt.Printf("ENTERED GETSHA: %s post",password)
	if ok {
		fmt.Printf("GET SHA password found: %s\n", password)
		return password, nil

	} else {
		fmt.Printf("GET SHA Password not found!\n")
		return "", errors.New("password does not exist")

	}
}

//TODO
func GetMD5(hash string) (string, error) {
	//return "", errors.New("not implemented")
	password, ok := md5lookup[hash]
	//fmt.Printf("ENTERED GETMD5: %s post",password)
	if ok {
		fmt.Printf("GET MD5 password found: %s\n", password)
		return password, nil

	} else {
		fmt.Printf("GET MD5 Password not found!\n")
		return "", errors.New("password does not exist")

	}
}
