package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
     count := 0
   
	for n != 1 {
       if n == 0 || n < 0 {
            return count, errors.New("Invalid n value")
           continue
      }
      if n % 2 == 0 {
          n = n / 2
      } else if n % 2 == 1{
          n = (n * 3) + 1
      }
    count += 1
    }
    return count, nil
}
