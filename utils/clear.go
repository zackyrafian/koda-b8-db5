package utils

import "fmt"

func Clear() { 
  // Source - https://stackoverflow.com/a/22892171
  // Posted by Kavu, modified by community. See post 'Timeline' for change history
  // Retrieved 2026-07-15, License - CC BY-SA 4.0
  
  fmt.Print("\033[H\033[2J")

}