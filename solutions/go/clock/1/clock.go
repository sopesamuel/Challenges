package clock
import "fmt"

// Define the Clock type here.
type Clock struct {
    Hour  int
    Minute int
}

func New(h, m int) Clock {
    newTime := h * 60 + m
    sumTime := newTime % 1440
    if sumTime < 0 {
        sumTime += 1440
    }
    newMinute := sumTime % 60
    newHour := sumTime / 60
    
    return Clock{Hour: newHour, Minute : newMinute}
}

func (c Clock) Add(m int) Clock {
    totalMinute := c.Hour * 60 + c.Minute 
    sumMinute := totalMinute + m
    sumMinute = sumMinute % 1440

  	if sumMinute < 0 {
        sumMinute += 1440
    }

    newH := sumMinute / 60
    newM := sumMinute % 60
     
	return Clock{Hour : newH, Minute: newM}
    
}

func (c Clock) Subtract(m int) Clock {
    totalMinute := c.Hour * 60 + c.Minute 
    sumMinute := totalMinute - m
    sumMinute = sumMinute % 1440

  	if sumMinute < 0 {
        sumMinute += 1440
    }

    newH := sumMinute / 60
    newM := sumMinute % 60
     
	return Clock{Hour : newH, Minute: newM}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d",c.Hour, c.Minute)
}
