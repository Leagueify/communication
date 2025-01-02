package server

import "fmt"

type Server interface {
	Start()
}

func showStartBanner() {
	fmt.Println(`
'||'      '||''''|      |      ..|'''.|  '||'  '|' '||''''|  '||' '||''''| '||' '|'
 ||        ||  .       |||    .|'     '   ||    |   ||  .     ||   ||  .     || |
 ||        ||''|      |  ||   ||    ....  ||    |   ||''|     ||   ||''|      ||
 ||        ||        .''''|.  '|.    ||   ||    |   ||        ||   ||         ||
.||.....| .||.....| .|.  .||.  ''|...'|    '|..'   .||.....| .||. .||.       .||.

  ..|'''.|  ..|''||   '||    ||' '||    ||' '||'  '|' '|.   '|' '||'   ..|'''.|     |     |''||''| '||'  ..|''||   '|.   '|'
.|'     '  .|'    ||   |||  |||   |||  |||   ||    |   |'|   |   ||  .|'     '     |||       ||     ||  .|'    ||   |'|   |
||         ||      ||  |'|..'||   |'|..'||   ||    |   | '|. |   ||  ||           |  ||      ||     ||  ||      ||  | '|. |
'|.      . '|.     ||  | '|' ||   | '|' ||   ||    |   |   |||   ||  '|.      .  .''''|.     ||     ||  '|.     ||  |   |||
 ''|....'   ''|...|'  .|. | .||. .|. | .||.   '|..'   .|.   '|  .||.  ''|....'  .|.  .||.   .||.   .||.  ''|...|'  .|.   '|
	`)
}
