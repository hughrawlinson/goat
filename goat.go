package goat

func zcr (buffer []float) int{
     int zcr = 0
     for i := 0; i < len(buffer); i++ {
     	 if (buffer[i] >= 0 && buffer[i+1] < 0) || (buffer[i] < 0 && buffer[i+1] >=0){
	     zcr++
	 }
     }
     return zcr
}
