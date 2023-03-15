package cmd

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var rootContext struct {
	hash string
}

func hashMethod(hashName string) hash.Hash {
	if hashName == "sha256" {
		return sha256.New()

	} else if hashName == "md5" {
		return md5.New()
	} else {
		fmt.Println("哈希算法不支持: " + hashName)
		return nil
	}
}

var rootCmd = &cobra.Command{
	Use:   "vhash",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("请输入哈希次数执行:")
		text, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		count, err := strconv.ParseInt(string(text), 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if count < 1 {
			fmt.Println("哈希次数必须大于0")
			return
		}
		fmt.Println("请输入文本:")
		text, err = term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for i := int64(0); i < count; i++ {
			alg := hashMethod(rootContext.hash)
			if alg == nil {
				return
			}
			_, err = alg.Write(text)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			h := alg.Sum(nil)
			text = []byte(hex.EncodeToString(h))
		}
		fmt.Println(string(text))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&rootContext.hash, "hash", "md5", "使用的哈希算法(sha256, md5)")
}
