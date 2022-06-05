package logics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	mg "vsmsforchild/msg"

	"github.com/asaskevich/govalidator"
)

var AnsYN, nameChild string
var VarConsent string
var AddRespVarMain []int
var Pdata []int
var ChildAge, ParentsAge int
var Cdata []int
var IsParent int
var IsChild int
var IsStart bool = false

//check expire time
func IsExpired() bool {
	//defer wg.Done()
	const expTime int = 190
	runingTime := time.Now().YearDay()

	return runingTime > expTime
}

//TTS with print
func TTSwithPrint(anyString string) {
	fmt.Println(">>  ", anyString)
	if anyString != "" {
		cmd := exec.Command(`./espk/espeak`, "-v", "f4", "-s", "140", "-p", "90", "-a", "100", anyString)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	} else {
		cmd := exec.Command(`./espk/espeak`, "-v", "f4", "-s", "140", "-p", "90", "-a", "100", "Sorry Wrong Input.")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

}

//TTS
func TTS(anyString string) {
	if anyString != "" {
		cmd := exec.Command(`./espk/espeak`, "-v", "f4", "-s", "140", "-p", "90", "-a", "100", anyString)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	} else {
		cmd := exec.Command(`./espk/espeak`, "-v", "f4", "-s", "140", "-p", "90", "-a", "100", "Sorry Wrong Input.")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

}

//screen clean
func ScreenCls() {
	cmd := exec.Command("powershell", "cls")

	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))

}

//check condition  in response as yes or no by user
func CheckConsent() bool {

	userInput := bufio.NewScanner(os.Stdin)

	for {
		//label:
		TTSwithPrint(strings.TrimLeft(mg.MsgHead[1], ""))
		fmt.Print(">>   ")
		time.Sleep(50 * time.Millisecond)
		userInput.Scan()
		AnsYN = strings.TrimSpace(userInput.Text())
		if AnsYN == "y" || AnsYN == "Y" {
			TTSwithPrint("Carefully read the to ask questions then answer.")
			return true
		} else if AnsYN == "n" || AnsYN == "N" {
			TTSwithPrint("Thanks For Visit......")
			time.Sleep(2000 * time.Millisecond)
			os.Exit(0)

		} else {
			fmt.Println(">>   Enter the correct reply Y or N.")
			time.Sleep(1500 * time.Millisecond)
			continue
			//TTSwithPrint("Enter the correct reply Y or N.")

			//goto label
		}

	}

}

//here checking the parents information
func CheckParent(str string) (is_parentAns int, err error) {
	var ans int

	userInput := bufio.NewScanner(os.Stdin)
	for {
		//label:
		TTSwithPrint(str)
		fmt.Print(">>   ")
		userInput.Scan()
		var in = strings.TrimSpace(userInput.Text())
		TTS("You entered the  " + in)
		if in == "y" || in == "Y" {
			ans = 1
			return ans, nil
		} else if in == "n" || in == "N" {
			return 0, nil
		} else {
			//TTSwithPrint("Enter the correct reply Y or N. ")
			fmt.Println(">>   Enter the correct reply Y or N.")
			//goto label
			continue
		}
	}

}

//all type age validation
func AgeValidator(s string) int {
	if govalidator.IsInt(s) {
		a, _ := strconv.Atoi(s)
		return a
	} else {
		//TTSwithPrint("Enter integer only")
		fmt.Println(">>   Enter integer only.")
		return 0
	}

}

//age validation uses in parents age
func AgeValidatorParent(age int) int {
	if age >= 22 && age <= 90 {
		return age
	} else if age < 22 && age > 90 {
		//TTSwithPrint("You are not came under 22 to 90 years. Try is by other parents.Bye...")
		fmt.Println(">>   You are not came under 22 to 90 years. Try is by other parents.Bye...")
		time.Sleep(2000 * time.Microsecond)
		os.Exit(0)
	}
	return 0
}

//age validation uses in child age
func AgeValidatorChild(age int) int {
	if age > 15 || age < 0 {
		//TTSwithPrint("This test for only under the 15 years children.")
		fmt.Println(">>   This test for only under the 15 years children.")
		return 0
	}
	return age
}

//answers takes in int value in 1 (yes) or 0  (no) and append in variable
func AnswerYNint() int {
	stndIo := bufio.NewReader(os.Stdin)
	//TTSwithPrint("Please enter  Y (Yes) or N (No).")
	for {
		//Loop:
		fmt.Print(">>   ")
		_, err := fmt.Fscan(stndIo, &VarConsent)
		if err != nil {
			TTS("Sorry no got any answer")
			panic(">>  Sorry no got any answer !...")
		}
		stndIo.ReadString('\n')
		get := strings.TrimLeft(VarConsent, "")
		TTS("You entered the " + get)
		if get == "Y" || get == "y" {
			return 1
		} else if get == "N" || get == "n" {

			return 0
		} else {
			//if VarConsent != "" {
			//TTSwithPrint("Sorry, invalid input please enter  Y (Yes) or N (No).")
			fmt.Println(">>   Sorry, invalid input please enter  Y (Yes) or N (No).")
			//}
			//VarConsent = ""
			//goto Loop
			continue
		}
	}
}

//information of parent getting here
func InfoParent() (doneInfo bool, respOfParent []int, child_name string, disabilityOfChild int) {
	var name_Child string
	var RespVar []int
	var disability int
	TTSwithPrint("What is the name of your child ?")
	fmt.Print(">>   ")
	fmt.Scanf("%s", &name_Child)
	TTS(nameChild)
	for i := 0; i < 4; i++ {
		TTSwithPrint(mg.GetQparents[i])
		RespVar = append(RespVar, AnswerYNint())
		if RespVar[0] == 1 && i == 0 {
			RespVar = append(RespVar, 0)
			i++
		}
		if i == 2 {
			disability = RespVar[2]
		}
	}

	return true, RespVar, name_Child, disability
}

//to get information of parent
func GetDetailsParents(is_parent, is_child int) (nowStart bool, parentsage int, childage int, responseOfParents []int, name_child string, disabilityOfChild int) {
	var age string
	var parents_age, child_age int
	var askingAge = []string{
		"Enter the age in between 22 to 90 years.",
		"Enter the age in between 0 to 15 years.",
	}
	if is_parent == 1 && is_child == 0 {

		stndIo := bufio.NewReader(os.Stdin)
		for i := 0; i < 2; i++ {
			TTSwithPrint(mg.GetQagePC[i])
		Loop1:
			TTSwithPrint(askingAge[i])
			fmt.Print(">>   ")
			_, err := fmt.Fscan(stndIo, &age)
			if err != nil {
				TTS("Sorry no got any answer")
				panic(">>  Sorry no got any answer !...")
			}
			TTS("You entered the  " + age)
			if i == 0 && IsParent != 0 {
				YesValid := AgeValidator(age)
				parents_age = AgeValidatorParent(YesValid)

				if YesValid != 0 && parents_age != 0 {
					age = "0"
				} else {
					//continue
					goto Loop1
				}
			}
			if i == 1 && IsChild == 0 {
				YesValid := AgeValidator(age)
				child_age = AgeValidatorChild(YesValid)

				if child_age == 0 {
					goto Loop1
					//continue
				}
				TTSwithPrint("Let starting the test.")
				start, respOfP, child_name, disability_child := InfoParent()
				IsStart = start
				return true, parents_age, child_age, respOfP, child_name, disability_child

			}
		}

	}
	return false, 0, 0, nil, "", 0
}

//to check the child is vaid or not
func CheckChild(s string) (doneInfo int, err error) {
	var ans int

	userInput := bufio.NewScanner(os.Stdin)
	for {
		//label:

		TTSwithPrint(mg.QIsChild)
		fmt.Print(">>   ")
		userInput.Scan()
		var in = userInput.Text()
		in = strings.TrimSpace(in)
		TTS("You entered the  " + in)
		if in == "y" || in == "Y" {
			ans = 1
			return ans, nil
		} else if in == "n" || in == "N" {
			return 0, nil
		} else {
			//TTSwithPrint("Enter the correct reply Y or N.")
			fmt.Println(">>   Enter the correct reply Y or N.")
			//goto label
			continue
		}
	}

}

//get information of child before start the test
func InfoChild() (doneInfo bool, response []int, nameofChild string, disabilyOfChild int) {
	var name_child string
	var RespVar []int
	var disability int
	TTSwithPrint("What is your name ?")
	fmt.Print(">>   ")
	fmt.Scanf("%s", &name_child)
	TTS(name_child)
	for i := 0; i < 4; i++ {
		TTSwithPrint(mg.GetQchild[i])
		RespVar = append(RespVar, AnswerYNint())
		if RespVar[0] == 1 && i == 0 {
			RespVar = append(RespVar, 0)
			i++
		}
		if i == 2 {
			disability = RespVar[2]
		}
	}
	return true, RespVar, name_child, disability
}

//taking the general information of child
func GetDetailsChild() (doneInfo bool, age_c int, respOfchild []int, nameChild string, disabilityOfChild int) {
	var age string
	var cAge int
	if IsParent == 0 && IsChild == 1 {
		TTSwithPrint(mg.GetQageC)
		for {

			//Loop1:

			stndIo := bufio.NewReader(os.Stdin)
			TTSwithPrint("Please  enter the age in between 0 to 15 years.")
			fmt.Print(">>   ")
			_, err := fmt.Fscan(stndIo, &age)
			if err != nil {
				TTS("Sorry no got any answer")
				panic(">>  Sorry no got any answer !...")
			}
			TTS("You entered the " + age)
			YesValid := AgeValidator(age)
			ChildAge = AgeValidatorChild(YesValid)
			cAge = AgeValidatorChild(YesValid)
			if ChildAge == 0 {
				//goto Loop1
				continue
			}
			TTSwithPrint("Let Starting the test.")
			start, resp, name_child, disability_child := InfoChild()
			IsStart = start
			return true, cAge, resp, name_child, disability_child
		}

	}

	return false, 0, nil, "", 0
}

func NewTestBegin() (reply bool) {
	fmt.Println("Do You Want Continue For Other Child ? Reply in Y or N.")
	var AnsYN string
	for {

		stdinput := bufio.NewReader(os.Stdin)
		fmt.Print(">>   ")
		_, err := fmt.Fscan(stdinput, &AnsYN)
		stdinput.ReadString('\n')
		if err != nil {
			fmt.Println("Thanks For Visit.....")
			//time.Sleep(3000 * time.Millisecond)
			os.Exit(0)

		} else if AnsYN == "y" || AnsYN == "Y" {
			for k := 0; k < 1; k++ {

				fmt.Println("Just Wait The Test Is Starting")
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("..")
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("...")
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("....\n\n")

			}
			fmt.Println("Just Wait The Test Is Starting")
			return true
		} else if AnsYN == "n" || AnsYN == "N" {
			fmt.Println("Thanks For Visit......")
			ScreenCls()
			//time.Sleep(3000 * time.Millisecond)
			os.Exit(0)

		} else {
			fmt.Println("Enter the correct reply Y or N.")
		}
	}
}
