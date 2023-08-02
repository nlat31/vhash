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
	hash    string
	show    bool
	reverse bool
	max     int64
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
		if !rootContext.reverse {
			fmt.Println("请输入哈希次数执行:")
			text, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if rootContext.show {
				fmt.Println(string(text))
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
			if rootContext.show {
				fmt.Println(string(text))
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
		} else {
			fmt.Println("请输入文本:")
			text, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if rootContext.show {
				fmt.Println(string(text))
			}
			fmt.Println("请输入哈希:")
			hashVal, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if rootContext.show {
				fmt.Println(string(hashVal))
			}
			times := int64(1)
			for ; times < rootContext.max; times++ {
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
				hashStr := hex.EncodeToString(h)
				if hashStr == string(hashVal) {
					break
				}
				text = []byte(hashStr)
			}
			fmt.Println(times)
		}
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
	rootCmd.Flags().BoolVar(&rootContext.show, "show", false, "显示输入内容")
	rootCmd.Flags().BoolVarP(&rootContext.reverse, "reverse", "r", false, "反向计算哈希次数")
	rootCmd.Flags().Int64VarP(&rootContext.max, "max", "m", 1000000, "反向计算哈希时的最大值")
}
