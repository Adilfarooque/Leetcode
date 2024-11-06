package main



//Move All zero into the end
func MoveAllzero(nums []int){
	nonZero := 0
    for i := 0; i < len(nums); i++ {
        nums[nonZero],nums[i] = nums[i], nums[nonZero]
        nonZero++
    }
}
func main(){
	
}