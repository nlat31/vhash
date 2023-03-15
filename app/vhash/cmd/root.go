package cmd

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var rootContext struct {
	hash  string
	count int
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
	Use:   "pwdgen",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if rootContext.count < 1 {
			fmt.Println("哈希次数必须大于1")
			return
		}
		fmt.Println("请输入文本:")
		text, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for i := 0; i < rootContext.count; i++ {
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
	rootCmd.Flags().IntVarP(&rootContext.count, "count", "c", 1, "哈希执行次数")
}
