package qans

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"time"

	ex "vsmsforchild/expitems"
	l "vsmsforchild/logics"
	"vsmsforchild/msg"
)

var AgeOfChild int
var BackTestData, CurrentTestData, ForwardTestData []float32
var AnswerOfQ string

type StructSliceChildInfo struct {
	ChildAgeSlice []int
	SumSlice      []float32
	SumAnsSlice   [][]float32
}

var SliceChildInfo StructSliceChildInfo

//start main test
func AskQuesTest(child_age int, disabily int) (isfinishQ bool, qDataStruct StructSliceChildInfo, frdCount []int, bckCount []int) {
	//defer wg.Done()
	var backCounter, frdCounter []int

	var ageSclice []int
	var arrSliceChild StructSliceChildInfo
	var age = child_age
	isdisability := disabily
	if age == 0 || isdisability == 1 {
		age = 1
	}
	for {
		backCounter = append(backCounter, age)
		frdCounter = append(frdCounter, age)
	Loop1:
		TestDone, Result := AgeSendSwitch(age)
		ageSclice = append(ageSclice, age)
		arrSliceChild.ChildAgeSlice = append(arrSliceChild.ChildAgeSlice, age)
		//fmt.Printf("TestDone:%+v,Result; %+v\n", TestDone, Result)
		ans := SumResult(Result)
		//fmt.Println("ans value ", ans)
		arrSliceChild.SumSlice = append(arrSliceChild.SumSlice, ans)
		arrSliceChild.SumAnsSlice = append(arrSliceChild.SumAnsSlice, Result)
		if TestDone {
			//fmt.Println("Current test is finished.", arrSliceChild)
			Qsize := float32(len(Result))
			time.Sleep(500 * time.Millisecond)
			if (ans >= (Qsize/100)*75 || ans == Qsize) && age < 13 {
				if age == 0 {
					age++
				}
				age++
				_, _, yesAgeF := AgeMatch(age, frdCounter)
				_, _, yesAgeB := AgeMatch(age, backCounter)
				_, _, yesAgeC := AgeMatch(age, ageSclice)
				time.Sleep(500 * time.Millisecond)
				if yesAgeF || yesAgeB || yesAgeC {
					break
				}
				for k := 0; k < 1; k++ {
					fmt.Printf(">>   ")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("Next ")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("Test ")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("Beginning ")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("..")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("...")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("....\n\n")

				}

				frdCounter = append(frdCounter, age)
				TestDone = false
				Result = nil
				goto Loop1
			} else if (Qsize/2) >= ans && age > 1 {
				age--
				_, _, yesAgeF := AgeMatch(age, frdCounter)
				_, _, yesAgeB := AgeMatch(age, backCounter)
				_, _, yesAgeC := AgeMatch(age, ageSclice)
				time.Sleep(500 * time.Millisecond)
				if yesAgeF || yesAgeB || yesAgeC {
					break
				}
				for k := 0; k < 1; k++ {
					fmt.Printf(">>   ")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("Previous ")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("Test ")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("Beginning ")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("..")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("...")
					time.Sleep(500 * time.Millisecond)
					fmt.Printf("....\n\n")

				}
				//fmt.Println(">>  Back Test Begin...")
				backCounter = append(backCounter, age)
				time.Sleep(500 * time.Millisecond)
				TestDone = false
				Result = nil
				goto Loop1
			}
			break
		}
		break
	}

	//fmt.Println("slice arr Slice ,frdcount, backcount,sliceage", arrSliceChild, frdCounter, backCounter, ageSclice)

	if arrSliceChild.SumSlice != nil {
		return true, arrSliceChild, frdCounter, backCounter
	}

	return false, StructSliceChildInfo{}, nil, nil
}

func AgeSendSwitch(AgeOfChild int) (bool, []float32) {
	switch AgeOfChild {
	case 0, 1:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.ZeroTo1[0:])
			return ChildAns, IsFinish
		}
	case 2:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.OneTo2[0:])
			return ChildAns, IsFinish
		}
	case 3:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.Twoto3[0:])
			return ChildAns, IsFinish
		}
	case 4:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.ThreeTo4[0:])
			return ChildAns, IsFinish
		}
	case 5:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.FourTo5[0:])
			return ChildAns, IsFinish
		}
	case 6:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.FiveTo6[0:])
			return ChildAns, IsFinish
		}
	case 7:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.SixTo7[0:])
			return ChildAns, IsFinish
		}
	case 8:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.SevenTo8[0:])
			return ChildAns, IsFinish
		}
	case 9:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.EightTo9[0:])
			return ChildAns, IsFinish
		}
	case 10:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.NineTo10[0:])
			return ChildAns, IsFinish
		}
	case 11:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.TenTo11[0:])
			return ChildAns, IsFinish
		}
	case 12:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.ElevenTo12[0:])
			return ChildAns, IsFinish
		}
	case 13, 14, 15:
		{
			ChildAns, IsFinish := TestByAgeOfChild(msg.TwelveTo15[0:])
			return ChildAns, IsFinish
		}
	default:
		l.TTSwithPrint("Age Not Under defined....")
		return false, nil
	}
}

func TestByAgeOfChild(TestName []string) (bool, []float32) {
	var currentTestData []float32
	q := 0
	l.TTSwithPrint("Enter Y (Yes) or N (No) or S (Sometime/By help other).")
	for q < len(TestName) {
	Loop:
		l.TTSwithPrint(TestName[q])
		Ans, _ := AnsOfQ()
		currentTestData = append(currentTestData, Ans)
		if currentTestData == nil {
			log.Println(">>  ", currentTestData)
			goto Loop
		}
		time.Sleep(1000 * time.Millisecond)
		q++
	}
	return true, currentTestData
}

func AnsOfQ() (float32, bool) {
	stndIo := bufio.NewReader(os.Stdin)
	for {
		//l.TTSwithPrint("Enter Y (Yes) or N (No) or S (Sometime/By help other).")
	Loop:
		fmt.Print(">>   ")
		_, err := fmt.Fscan(stndIo, &AnswerOfQ)
		if err != nil {
			l.TTS("Sorry no got any answer")
			panic(">>  Sorry no got any answer !...")
		}
		stndIo.ReadString('\n')
		get := strings.TrimLeft(AnswerOfQ, "")
		l.TTS("You entered the " + get)
		if get == "Y" || get == "y" {
			return 1.0, true
		} else if get == "N" || get == "n" {
			return 0.0, true
		} else if get == "S" || get == "s" {
			return 0.5, true
		} else {
			if AnswerOfQ != "" {
				l.TTSwithPrint("Sorry, Invalid input. Please enter Y (Yes)/ N (No)/ S (Sometime/By help other).")
			}
			AnswerOfQ = ""
			goto Loop
		}
	}
}

func AnswerMatch(findVal int, arr []int) (idx int, val int, b bool) {
	i := 0
	for i < len(arr) {
		if findVal == arr[i] {
			return i, findVal, true
		}
		i++
	}
	return 0, findVal, false
}

func AgeMatch(findVal int, arr []int) (idAge int, val int, b bool) {
	i := 0
	for i < len(arr) {
		if findVal == arr[i] {
			return i, findVal, true
		}
		i++
	}
	return 0, findVal, false
}
func SumResult(f []float32) float32 {
	var a float32
	for _, val := range f {
		a += val
	}
	return a
}

func ValidQuestionFind(age int, ans []float32) (ageNum int, questionNum []int, answerNum []float32) {

	if ans != nil {

		switch age {

		case 0, 1:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(1, 17, age, ans)
				return ageNum, quesNum, ansNum
			}
		case 2:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(18, 34, age, ans)
				return ageNum, quesNum, ansNum
			}
		case 3:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(35, 44, age, ans)
				return ageNum, quesNum, ansNum
			}
		case 4:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(45, 50, age, ans)
				return ageNum, quesNum, ansNum
			}
		case 5:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(51, 56, age, ans)
				return ageNum, quesNum, ansNum
			}
		case 6:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(57, 61, age, ans)
				return ageNum, quesNum, ansNum
			}
		case 7:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(62, 65, age, ans)
				return ageNum, quesNum, ansNum
			}
		case 8:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(66, 70, age, ans)
				return ageNum, quesNum, ansNum
			}
		case 9:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(71, 74, age, ans)
				return ageNum, quesNum, ansNum
			}
		case 10:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(75, 77, age, ans)
				return ageNum, quesNum, ansNum
			}
		case 11:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(78, 81, age, ans)
				return ageNum, quesNum, ansNum
			}
		case 12:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(82, 84, age, ans)
				return ageNum, quesNum, ansNum

			}
		case 13, 14, 15:
			{
				ageNum, quesNum, ansNum := ReturnQuestionNum(85, 89, age, ans)
				return ageNum, quesNum, ansNum
			}
		default:
			fmt.Println(">>  Not Under defined....")
			return age, nil, nil
		}
	}
	fmt.Printf(">>  Arguments Empty !  age %+v, ans %+v\n", age, ans)
	return 0, nil, nil
}

func ReturnQuestionNum(iIndex int, lIndex int, age int, ans []float32) (ageNum int, question []int, answer []float32) {
	var questionsNum []int
	var answerNum []float32
	j := 0
	var one float32 = 1.0
	var half float32 = 0.5
	if ans != nil {
		for i := iIndex; i <= lIndex; i++ {
			if ans[j] == one {
				//fmt.Printf("ans::%+v", i)
				questionsNum = append(questionsNum, i)
				answerNum = append(answerNum, one)
			} else if ans[j] == half {
				//fmt.Printf("ans::%+v", i)
				questionsNum = append(questionsNum, i)
				answerNum = append(answerNum, half)
			}
			j++
		}
		return age, questionsNum, answerNum
	}
	fmt.Printf(">>  Arguments Empty ! iIndex %+v, lIndex %+v, age %+v, ans %+v\n", iIndex, lIndex, age, ans)
	return 0, nil, nil
}

func QuesBindMaxMinNum(arrQuestions []int) (minNum int, maxNum int, bindedQuestion []int) {
	var min int = arrQuestions[0]
	var max int = arrQuestions[0]
	var bindQues []int

	if arrQuestions != nil {
		for _, val := range arrQuestions {
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
			bindQues = append(bindQues, val)
		}
		return min, max, bindQues
	}
	fmt.Printf(">>  Error in argruments like this %+v\n", arrQuestions)
	return 0, 0, nil
}

func StaticsWork(sliceChild *StructSliceChildInfo) (IsDone bool, MinAge int, MaxAge int, BindQuestionsNum []int, SocialAge float32, SocialQuotient float32, SocialAgeInMonth float32, AgeInYear float64, AgeInMonth float64) {
	var preChildScore float32
	var gotChildScore float32
	var age_year, age_month float64
	var gotChildQuestion []int
	var rawScore, social_age, newSAinMonth float32

	//defer wg.Done()
	//previous answers add in variavle before minimum age
	minAge, maxAge, _ := QuesBindMaxMinNum(sort.IntSlice(sliceChild.ChildAgeSlice))

	//fmt.Printf("Min Age: %+v, Max Age: %+v,Bind Age: %+v", minAge, maxAge, bindAge)

	// for ages
	for i := 0; i < len(sliceChild.ChildAgeSlice); i++ {

		_, Q_NUM, ANS_NUM := ValidQuestionFind(sliceChild.ChildAgeSlice[i], sliceChild.SumAnsSlice[i])

		//fmt.Printf("Age: %+v, QuesNum: %+v,AnsNum: %+v\n", AGE, Q_NUM, ANS_NUM)
		gotChildQuestion = append(gotChildQuestion, Q_NUM...)
		//fmt.Printf("gotChildQuestion: %+v, gotChildScore: %+v\n", gotChildQuestion, gotChildScore)

		for j := 0; j < len(ANS_NUM); j++ {
			gotChildScore += ANS_NUM[j]
			//fmt.Printf("gotChildScore %+v\n", gotChildScore)
		}

	}
	sort.Ints(gotChildQuestion)
	fmt.Printf("####################################################################################################################################\n\n")
	fmt.Printf(">>   Child Score. : %v\n", gotChildScore)
	fmt.Printf(">>   Correct Attempted  Question. : %v\n", gotChildQuestion)
	minQ, _, _ := QuesBindMaxMinNum(gotChildQuestion)
	//fmt.Printf("minQ %v, maxQ  %v\n", minQ, maxQ)

	preChildScore = float32(minQ - 1)
	gotChildScore += preChildScore
	//fmt.Printf("gotChildScore %v, preChildScore  %v\n", gotChildScore, preChildScore)
	rawScore = float32(math.Round(float64(gotChildScore)))
	fmt.Printf(">>   Raw Score of Child. : %v \n", rawScore)

	// for key, val := range ex.RawScale {

	// }
	if val, found := ex.RawScale[rawScore]; found {
		social_age = val
		//fmt.Printf("found  %+v, val %+v\n", found, val)
		fmt.Printf(">>   Social Age of Child is. : %+v\n", social_age)
		yearVal, monthVal := math.Modf(float64(social_age))
		age_year = yearVal
		age_month = math.Round(monthVal * 10)
		fmt.Printf(">>   %+v Year and ", yearVal)
		fmt.Printf("%+v Month of Child Age.\n", math.Round(monthVal*10))
		SAinMonth := math.Round(float64(social_age * 12))
		newSAinMonth = float32(SAinMonth)
		//fmt.Printf("newSAinMonth: %+v\n", newSAinMonth)

	}

	SQ := (social_age / float32(sliceChild.ChildAgeSlice[0])) * 100
	S_quotient := float32(math.Round(float64(SQ)))
	fmt.Printf(">>   Social Quotient is. : %+v\n", S_quotient)
	//fmt.Printf("Social Age  %+v\n", social_age)
	fmt.Printf("####################################################################################################################################\n\n")
	if SQ != 0 {
		return true, minAge, maxAge, gotChildQuestion, social_age, S_quotient, newSAinMonth, age_year, age_month
	}
	return false, minAge, maxAge, gotChildQuestion, social_age, S_quotient, newSAinMonth, age_year, age_month
	//FindDataSA(newSAinMonth, &ex.NormsAnalysis, minQ, gotChildQuestion)
}

func ShowCharacteristics(SAinMonth float32, scalNorm *ex.ScaleNorm, minQuestionNum int, gotChildQuestion []int) (isDone bool, bindCharts map[string]string) {
	localScaleMonth := scalNorm.Months
	localScaleSHG := scalNorm.SHG
	localScaleSHE := scalNorm.SHE
	localScaleSHD := scalNorm.SHD
	localScaleSD := scalNorm.SD
	localScaleOCC := scalNorm.OCC
	localScaleCOM := scalNorm.COM
	localScaleLOC := scalNorm.LOC
	localScaleSOC := scalNorm.SOC
	var social_age float32 = SAinMonth
	//var lamsumFindSA float32 = float32(math.Round(float64(findSA)))
	bindCharacterstics := make(map[string]string)
	for i := 0; i < len(localScaleMonth); i++ {

		//fmt.Println("localScaleMonth[i]:", localScaleMonth[i])
		//fmt.Println("lamsumFindSA:", lamsumFindSA)

		for j := 0; j < len(localScaleMonth[i]); j++ {
		LOOP:
			if social_age == localScaleMonth[i][j] {
				fmt.Println(">>   Characteristics of child are/is showing at below.:-")
				//fmt.Println("findSA:", social_age, "localScaleMonth[i][j]", localScaleMonth[i][j], "index of search :", i, j)

				//use for Self Help  General (SHG)
				chartsForSHG, QforSHG := FindCharacteristic(i, j, localScaleSHG, gotChildQuestion)
				if chartsForSHG && QforSHG {
					//bindCharacterstics["Self Help  General:"] = "Yes"
					fmt.Println(">>   Self Help  General: Yes")
				} else {
					//bindCharacterstics["Self Help  General:"] = "No"
					fmt.Println(">>   Self Help  General: No")
				}

				//use for Self Help  Eating (SHE)
				chartsForSHE, QforSHE := FindCharacteristic(i, j, localScaleSHE, gotChildQuestion)
				if chartsForSHE && QforSHE {
					//bindCharacterstics["Self Help Eating:"] = "Yes"
					fmt.Println(">>   Self Help Eating: Yes")
				} else {
					//bindCharacterstics["Self Help Eating:"] = "No"
					fmt.Println(">>   Self Help Eating: No")
				}

				//use for Self Help Dressing   (SHD)
				chartsForSHD, QforSHD := FindCharacteristic(i, j, localScaleSHD, gotChildQuestion)
				if chartsForSHD && QforSHD {
					bindCharacterstics["Self Help Dressing:"] = "Yes"
					fmt.Println(">>   Self Help Dressing: Yes")
				} else {
					bindCharacterstics["Self Help Dressing:"] = "No"
					fmt.Println(">>   Self Help Dressing: No")
				}

				//use for Self-Direction   (SD)
				chartsForSD, QforSD := FindCharacteristic(i, j, localScaleSD, gotChildQuestion)
				if chartsForSD && QforSD {
					bindCharacterstics["Self-Direction:"] = "Yes"
					fmt.Println(">>   Self-Direction: Yes")
				} else {
					bindCharacterstics["Self-Direction:"] = "No"
					fmt.Println(">>   Self-Direction: No")
				}

				//use for Occupation   (OCC)
				chartsForOCC, QforOCC := FindCharacteristic(i, j, localScaleOCC, gotChildQuestion)
				if chartsForOCC && QforOCC {
					bindCharacterstics["Occupation:"] = "Yes"
					fmt.Println(">>   Occupation: Yes")
				} else {
					bindCharacterstics["Occupation:"] = "No"
					fmt.Println(">>   Occupation: No")
				}

				//use for Communication(COM)
				chartsForCOM, QforCOM := FindCharacteristic(i, j, localScaleCOM, gotChildQuestion)
				if chartsForCOM && QforCOM {
					bindCharacterstics["Communication:"] = "Yes"
					fmt.Println(">>   Communication: Yes")
				} else {
					bindCharacterstics["Communication:"] = "No"
					fmt.Println(">>   Communication: No")
				}

				//use for Locomotion(LOC)
				chartsForLOC, QforLOC := FindCharacteristic(i, j, localScaleLOC, gotChildQuestion)
				if chartsForLOC && QforLOC {
					bindCharacterstics["Locomotion:"] = "Yes"
					fmt.Println(">>   Locomotion: Yes")
				} else {
					bindCharacterstics["Locomotion:"] = "No"
					fmt.Println(">>   Locomotion: No")
				}

				//use for Socialisation(SOC)
				chartsForSOC, QforSOC := FindCharacteristic(i, j, localScaleSOC, gotChildQuestion)
				if chartsForSOC && QforSOC {
					bindCharacterstics["Socialisation:"] = "Yes"
					fmt.Println(">>   Socialisation: Yes")
				} else {
					bindCharacterstics["Socialisation:"] = "No"
					fmt.Println(">>   Socialisation: No")
				}

				fmt.Printf("##################################################################\n\n")
				//break
			} else if (social_age - 0.5) == localScaleMonth[i][j] {
				goto LOOP
			}

		}

	}
	//DefineAboutQ(ex.ItemExp, gotChildQuestion)
	//fmt.Printf("bindCharacterstics :%+v\n", bindCharacterstics)
	if len(bindCharacterstics) < 1 {

		return false, bindCharacterstics
	}
	return true, bindCharacterstics

}

func FindCharacteristic(indexI int, indexJ int, ArrayCharacteristics [][]int, qNums []int) (Characteristic bool, attempQuestion bool) {
	for _, valCHARTS := range ArrayCharacteristics[indexI] {
		if len(ArrayCharacteristics[indexI]) >= 1 && indexJ+1 <= len(ArrayCharacteristics[indexI]) {
			if ArrayCharacteristics[indexI][indexJ] != 0 {
				for _, findq := range qNums {
					if ArrayCharacteristics[indexI][indexJ] == findq {
						_ = valCHARTS //not use here
						//fmt.Println("localScaleSHD yes ::", "findq:", findq, "valSHD:", valCHARTS)
						return true, true
					}
				}
			}
		}

	}
	return false, false
}

//Summary of Test
func DefineAboutQ(itemsExpression map[int]string, qNums []int, socialAgeinMonth float32) (isDone bool, aboutQuestion []string, briefSocialAge string) {
	var bindAboutQuestion []string
	var defSocialAge string
	switch true {
	case (socialAgeinMonth < 20):
		{
			defSocialAge = "Profound Intellectual Disabiliy."

		}
	case (socialAgeinMonth < 34):
		{
			defSocialAge = "Severe Intellectual Disabiliy."

		}
	case (socialAgeinMonth < 49):
		{
			defSocialAge = "Moderate Intellectual Disabiliy."

		}
	case (socialAgeinMonth < 59):
		{
			defSocialAge = "Mild Intellectual Disabiliy."

		}
	case (socialAgeinMonth < 84):
		{
			defSocialAge = "Borderline Intellectual Disabiliy."

		}
	case (socialAgeinMonth < 110):
		{
			defSocialAge = "Average Developmental Functiong"

		}
	case (socialAgeinMonth < 129):
		{
			defSocialAge = "Above Developmental Functioning."

		}
	case (socialAgeinMonth < 150 || socialAgeinMonth >= 150):
		{
			defSocialAge = "Superior Developmental Functioning."

		}
	}

	for i := 0; i < len(qNums); i++ {
		questionNum := qNums[i]
		for key, val := range itemsExpression {
			if key == questionNum {
				bindAboutQuestion = append(bindAboutQuestion, val)
				//fmt.Println(">>  ", val)
				break
			}
		}
	}
	if bindAboutQuestion != nil {
		return true, bindAboutQuestion, defSocialAge
	}
	return false, bindAboutQuestion, defSocialAge
}
