package main

import (
	"crypto/rand"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	"math/big"
	"sync"
)

func main() {
	db , err := sql.Open("mysql", "root:somepass@(127.0.0.1:3306)/nimoc_blog") ; if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SetMaxIdleConns(40)
	db.SetMaxOpenConns(40)
	kinds := [...]string{"wait", "done", "fail"}
	thousand := 1000
	million := thousand * thousand
	wg := sync.WaitGroup{}
	for i:=0;i<40;i++{
		wg.Add(1)
		go func() {
			unit := 10*million
			for j:=0;j<unit;i++ {
				_, err := db.Exec("INSERT INTO `huge_data` (`id`, `title`, `amount`, `kind`) VALUES (?,?,?,?)", uuid.New(), BytesLetter(10), IntRange(0,10000000), kinds[IntRange(0, len(kinds)-1)] ) ; if err != nil {
					panic(err)
				}
			}
			log.Printf("created %s", i * unit)
			wg.Done()
		}()
	}
	wg.Wait()
}

func IntRange(min int, max int) int {
	if max == min {return max}
	value, err :=rand.Int(rand.Reader, big.NewInt(int64(max-min+1))) ; if err !=nil {panic(err)}
	return int(value.Int64())+min
}
func BytesBySeed(seed []byte, size int) []byte {
	result := []byte("")
	for i:=0; i<size; i++ {
		randIndex, err :=rand.Int(rand.Reader, big.NewInt(int64(len(seed)))) ; if err !=nil {panic(err)}
		result = append(result, seed[randIndex.Int64()])
	}
	return result
}
func BytesLetter (size int) []byte {
	return BytesBySeed([]byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a}, size)
}