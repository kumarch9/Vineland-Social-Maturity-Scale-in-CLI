package msg

var MsgHead = [2]string{
	"This test based on social maturity of children (VSMS Indain adapted) and it is a psychological test  under the age of 15 years " +
		"and this test can conduct by self, children's parents or any psychologist. Please keep in mind that all will be replied in honesty and true the asked  questions by " +
		"self or parents.",

	"Are you ready to conduct this test, Y (Yes) or N (No) ?",
}

//FAQ's from Parents
var QIsParents = "Are you parents of child ?" //if  response no, exit from it
var GetQagePC = [2]string{
	"What is your age ?",
	"What is your child age ?", // all will asking questions depend on it
}
var GetQparents = [4]string{
	"Your child is male ?",
	"Your child is female ?",
	"Your child is suffer any disability abled", //if yes,exit from it
	"Sure are you ready to conduct this test ?", //if yes, play it
}

//FAQ's from child
var QIsChild = "Are you a child and under in the age of 15 years ?" //if response no, exit from it
var GetQageC = "What is your age ?"                                 // all will asking questions depend on it
var GetQchild = [4]string{
	"You are  male ?",
	"You are  female ?",
	"You are suffer any disability abled",       //if yes,exit from it
	"Sure are you ready to conduct this test ?", //if yes, play it
}

//questions for 0 to 1
var ZeroTo1 = [17]string{
	"1. Crows, Laugh ?",
	"2. Balance head ?",
	"3. Grasps objects within reach ?",
	"4. Reaches for familiar persons ?",
	"5. Rolls over, (unassisted) ?",
	"6. Reaches for nearby objects ?",
	"7. Occupies self-upright ?",
	"8. Sits unsupported ?",
	"9. Pulls self upright ?",
	"10. Talks”, imitates sounds ?",
	"11. Drinks from cup or glass assisted ?",
	"12. Moves about on floor (creeping, crawling) ?",
	"13. Grasps with thumb and finger  ?",
	"14. Demands personal attention ?",
	"15. Stands alone ?",
	"16. Does not drool ?",
	"17. Follows simple instructions ?",
}

//questions for 1 to 2
var OneTo2 = [17]string{
	"18. Walks about room unattended ?",
	"19. Marks with pencil or crayon or chalk ?",
	"20. Masticates (chews) solid or semi-solid food ?",
	"21. Pulls off clothes ?",
	"22. Transfers objects ?",
	"23. Overcomes simple obstacles ?",
	"24. Fetches or carries familiar objects ?",
	"25. Drinks from cup or glass ?",
	"26. Walks without support ?",
	"27. Plays with other children ?",
	"28. Eats with own hands (biscuits, bread, etc.)  ?",
	"29. Goes about hours or yard ?",
	"30. Discriminates edible substances from non-edibles  ?",
	"31. Uses names of familiar objects ?",
	"32. Walks upstairs unassisted ?",
	"33. Unwraps sweets, chocolates ?",
	"34. Talks in short sentences ?",
}

//questions for 2 to 3
var Twoto3 = [10]string{
	"35. Signals to go to toilet ?",
	"36. Initiates own play activities ?",
	"37. Removes shirt or frock if unbuttond ?",
	"38. Eats with spoon/hands (food) ?",
	"39. Gets drink (water) unassisted ?",
	"40. Dries own hands ?",
	"41. Avoids simple hazards ?",
	"42. Puts on short or frock unassisted (need not button) ?",
	"43. Can do paper folding ?",
	"44. Relates experience ?",
}

//questions for 3 to 4
var ThreeTo4 = [6]string{
	"45. Walks downstairs, one step at a time ?",
	"46. Plays co-operatively at kindergarten level ?",
	"47. Buttons shirt or frock ?",
	"48. Helps at little household tasks ?",
	"49. “Performs” for others ?",
	"50. Washes hands unaided ?",
}

//questions for 4 to 5
var FourTo5 = [6]string{
	"51. Cares for self at toilet ?",
	"52. Washes face unassisted ?",
	"53. Goes about neighbourhood unattended ?",
	"54. Dresses self except for trying ?",
	"55. Uses pencil or crayon or chalk for drawing ?",
	"56. Plays competitive exercise games ?",
}

////questions for 5 to 6
var FiveTo6 = [5]string{
	"57. Uses hoops, flies kites,or uses knife ?",
	"58. Prints (writes) simple words ?",
	"59. Plays simple games which require talking turns ?",
	"60. Is trusted with money ?",
	"61. Goes to school unattended ?",
}

var SixTo7 = [4]string{
	"62. Mixex rice “properly unassisted ?",
	"63. Uses pencil or chalk for waiting ?",
	"64. Bathes self assisted ?",
	"65. Goes to bed unassisted ?",
}

//questions for 7 to 8
var SevenTo8 = [5]string{
	"66. Can differentiate betweween AM & PM ?",
	"67. Helps himself during meals ?",
	"68. Understands and keeps family secrets ?",
	"69. Participants in pre-adolescent ?",
	"70. Combs or bruses hair ?",
}

//questions for 8 to 9
var EightTo9 = [4]string{
	"71. Uses tools or utensils ?",
	"72. Does routine household tasks ?",
	"73. Reads on own initiative ?",
	"74. Bathes self unaided ?",
}

//questions for 9 to 10
var NineTo10 = [3]string{
	"75. Cares for self at meals ?",
	"76. Makes minor purchase ?",
	"77. Goes about home town freely ?",
}

//questions for 10 to 11
var TenTo11 = [4]string{
	"78. Distinguishes between friends any play mates ?",
	"79. Makes independent choice of shops ?",
	"80. Does small remunerative work makes articles ?",
	"81. Follows local current events ?",
}

var ElevenTo12 = [3]string{
	"82. Does simple creative work ?",
	"83. Is left to care for self or others ?",
	"84. Enjoys reading books, newspapers and magazines ?",
}

//questions for 12 to 15
var TwelveTo15 = [5]string{
	"85. Plays difficult games ?",
	"86. Exercises complete care of dress ?",
	"87. Buys own clothing accessories ?",
	"88. Engages of adolescent group activities ?",
	"89. Performs responsible routine chores ?",
}
