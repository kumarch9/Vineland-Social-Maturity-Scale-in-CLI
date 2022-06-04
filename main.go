package main

import (
	"fmt"
	"os"
	"sync"
	"time"
	ex "vsmsforchild/expitems"
	l "vsmsforchild/logics"
	mg "vsmsforchild/msg"
	q "vsmsforchild/qans"

	ct "github.com/daviddengcn/go-colortext"
)

var Disability, Childage, Parentsage, MaxQuesNum, MinQuesNum, isParent, isChild, minAge, maxAge int
var dataOfP, dataOfC, bindQuestionNumber, forwardCounter, backwardCounter []int

var varSliceChildInfo q.StructSliceChildInfo
var QuestionNum [][]int
var nameofChild, nameofParents, typeOfDisability string
var UserConsent bool
var IsDone []bool
var socialAge, socialAgeInMonth, socialQuotient float32

var AgeInYear, AgeInMonth float64

var bindCharacteristics map[string]string
var briefOfQNum []string
var wg sync.WaitGroup

func main() {

	//welcome banner

	ct.Foreground(ct.Blue, true)

	wg.Add(1)
	go getGeneralInfo(&wg)

	wg.Wait()

	//q.DefineAboutQ(ex.ItemExp, []int{57, 58, 59, 60, 61, 62, 63, 64}, 76)
}

//begin test function uses to get general information of child or parents
func getGeneralInfo(wg *sync.WaitGroup) {
	defer wg.Done()

	isExpTest := l.IsExpired()
	if isExpTest {
		fmt.Printf("VSMS Has Been Expired, Please Wait Test Is Exiting")
		time.Sleep(2000 * time.Millisecond)
		os.Exit(0)
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println(">>   ----WELCOME IN VSMS TEST----")
	for {
	ContinueNew:
		//vsmsBanner1 := fig.NewFigure("WELCOME IN VSMS\n\n\n", "", true)
		//vsmsBanner1.Print()
		//continue toward  test
		okConsent := l.CheckConsent()
		UserConsent = okConsent
		if !okConsent {
			l.TTSwithPrint("......Thanks visit gain, Bye......")
			os.Exit(0)
		}
		time.Sleep(2000 * time.Millisecond)

		OkParent, _ := l.CheckParent(mg.QIsParents)
		isParent = OkParent
		//fmt.Println("OkP:", OkParent)

		l.IsParent = OkParent
		//fmt.Println("isP and Isc", isParent, isChild)
		time.Sleep(1000 * time.Millisecond)
		IamParent, pAge, cAge, respOfP, childname, childDisability := l.GetDetailsParents(isParent, isChild)
		Parentsage = pAge
		Childage = cAge
		dataOfP = respOfP
		nameofChild = childname
		Disability = childDisability
		//fmt.Println("IamParent, pAge, cAge, respOfP, childname, childDisability", IamParent, pAge, cAge, respOfP, childname, childDisability)
		if !IamParent {
			l.TTSwithPrint("Parent detail is incomplete.")
		}

		if OkParent != 1 {
			okChild, _ := l.CheckChild(mg.QIsChild)

			isChild = okChild
			//fmt.Println("okC,", okChild)

			l.IsChild = okChild
			if okChild != 1 {
				l.TTSwithPrint("This test is not able for you,Thanks for visit.  Bye........")
				time.Sleep(2000 * time.Microsecond)
				os.Exit(0)
			}
			IamChild, childAge, respOfchild, nameChild, disabilityOfChild := l.GetDetailsChild()
			//fmt.Println("IamChild, childAge, respOfchild, nameChild, disabilityOfChild", IamChild, childAge, respOfchild, nameChild, disabilityOfChild)
			Childage = childAge
			dataOfC = respOfchild
			nameofChild = nameChild
			Disability = disabilityOfChild
			if !IamChild {
				l.TTSwithPrint("Incomplete details of child.")
			}

		}
		time.Sleep(1000 * time.Millisecond)
		l.ScreenCls()
		for k := 0; k < 1; k++ {

			fmt.Printf("Please ")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Wait, ")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Test ")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Is ")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Beginning")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("..")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("...")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("....\n\n")
		}

		is_finish, q_struct, frd_counter, bck_counter := q.AskQuesTest(Childage, Disability)
		if is_finish {
			IsDone = append(IsDone, is_finish)
			varSliceChildInfo = q_struct
			forwardCounter = frd_counter
			backwardCounter = bck_counter
			//fmt.Println("is_finish, q_struct, frd_counter, bck_counter", is_finish, q_struct, frd_counter, bck_counter)
			l.ScreenCls()
		}

		l.ScreenCls()
		for k := 0; k < 1; k++ {
			fmt.Printf("Please ")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Wait, ")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Report ")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Is ")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Generating ")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("..")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("...")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("....\n\n")
		}
		time.Sleep(1000 * time.Millisecond)
		l.ScreenCls()
		fmt.Println("---- Report is generated. ---- ")
		is_doneStatics, min_age, max_age, bind_qNum, social_age, social_quotient, social_ageInMonth, years_age, year_ageMonth := q.StaticsWork(&varSliceChildInfo)

		if is_doneStatics {
			IsDone = append(IsDone, is_doneStatics)
			minAge = min_age
			maxAge = max_age
			bindQuestionNumber = bind_qNum
			socialAge = social_age
			socialQuotient = social_quotient
			socialAgeInMonth = social_ageInMonth
			AgeInYear = years_age
			AgeInMonth = year_ageMonth
			//fmt.Println("is_doneStatics, min_age, max_age, bind_qNum, social_age, social_quotient, social_ageInMonth, years_age, year_ageMonth", is_doneStatics, min_age, max_age, bind_qNum, social_age, social_quotient, social_ageInMonth, years_age, year_ageMonth)
		}

		time.Sleep(1000 * time.Millisecond)
		isDone_charts, bind_charts := q.ShowCharacteristics(socialAgeInMonth, &ex.NormsAnalysis, MinQuesNum, bindQuestionNumber)
		if isDone_charts {
			IsDone = append(IsDone, isDone_charts)
			bindCharacteristics = bind_charts

			//fmt.Println("isDone_charts, bind_charts", isDone_charts, bind_charts)
		}

		time.Sleep(1000 * time.Millisecond)
		isDoneDefine, aboutQuestion, briefSA := q.DefineAboutQ(ex.ItemExp, bindQuestionNumber, socialAgeInMonth)
		if isDoneDefine {
			IsDone = append(IsDone, isDoneDefine)
			briefOfQNum = aboutQuestion
			typeOfDisability = briefSA

			fmt.Printf(">>   About Of %v :-\n", nameofChild)
			for _, val := range briefOfQNum {
				fmt.Println(">>   ", val)
			}
			fmt.Println(">>   Types of Disabilities : ", typeOfDisability)
			fmt.Printf("####################################################################################################################################\n\n")
		} else if !isDoneDefine {
			fmt.Println(">>   Types of Disabilities : ", briefSA)
			fmt.Printf("####################################################################################################################################\n\n")
		}
		time.Sleep(3000 * time.Millisecond)
		isTestStart := l.NewTestBegin()
		{
			if isTestStart {
				l.ScreenCls()
				goto ContinueNew
			}

			//here the compilar wil not come but on here the need to use of variable . if its wil any need in upoming to save file of information so it will use of variable.
			fmt.Println("minAge, maxAge, bindQuestionNumber, forwardCounter, backwardCounter, socialAge, socialQuotient, bindCharacteristics", minAge, maxAge, bindQuestionNumber, forwardCounter, backwardCounter, socialAge, socialQuotient, bindCharacteristics)
			fmt.Println("forwardCounter, isParent, isChild ", forwardCounter, isParent, isChild)
			fmt.Println("backwardCounter ..", backwardCounter)
			fmt.Println("varSliceChildInfo, dataOfP, dataOfC, nameofChild, nameofParents, socialAgeInMonth", varSliceChildInfo, dataOfP, dataOfC, nameofChild, nameofParents, socialAgeInMonth)

		}
	}

}
