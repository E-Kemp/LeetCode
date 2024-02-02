public class Solution {
    public IList<int> SequentialDigits(int low, int high) {
        string NUMS = "123456789";
        int lowNumDigits = low.ToString().Length;
        int highNumDigits = high.ToString().Length;
        IList<int> output = new List<int>();
        
        for(int numDigits = lowNumDigits; numDigits <= highNumDigits; numDigits++) {
            bool complete = false;

            // iterations = (length + 1) - size
            for(int index = 0; index < (NUMS.Length + 1) - numDigits && !complete; index++) {
                int num = Int32.Parse(NUMS.Substring(index, numDigits));
                if(low <= num && num <= high) {
                    output.Add(num);
                }
                else if(num > high) {
                    complete = true;
                }
            }
        }
        return output;
    }
}