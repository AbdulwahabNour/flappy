package main

import (
	"fmt"

	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const(
    FONT_DIR ="res/fonts/Flappy.ttf"
    BACKGROUND_DIR = "res/imgs/background.png"
    SCREEN_HIGHT = 800
    SCREEN_WIDTH = 1000 
)

func main(){
 
    if err := run(); err != nil{
         log.Fatalln(err)
    }
   
 
}
func run()error{
    err := sdl.Init(sdl.INIT_EVERYTHING)
    if err != nil{
      return fmt.Errorf("could not initialize SDL: %v", err)
    }
    defer sdl.Quit()
     
    if err :=ttf.Init(); err != nil{
        return fmt.Errorf("could not initialize ttf: %v", err)
    }
    defer ttf.Quit()
    w, r, err := sdl.CreateWindowAndRenderer(SCREEN_WIDTH, SCREEN_HIGHT, sdl.WINDOW_SHOWN)
    if err != nil{
        return fmt.Errorf("could not create window: %v", err)
    }
    defer w.Destroy()
    if err = drawTitle(r, "Welcome"); err != nil{
        return fmt.Errorf("could not draw title: %v", err)
    }
     
    time.Sleep(3 * time.Second)
    s, err := newScene(r)

    if err != nil{
        return fmt.Errorf("could not draw background: %v", err)
    }
    
 
    defer s.destroy()

    events := make(chan sdl.Event)
    e := s.runBird(r, events)
  
     for {
           select{
                   case  events <- sdl.WaitEvent():
                   case err := <-e:
                   
                    return err


           }
     }
    

 
}

func drawTitle(r *sdl.Renderer, msg string) error{
    f, err := ttf.OpenFont(FONT_DIR, 20)
    if err != nil{
         return fmt.Errorf("could not open font %v", err)
    }
    defer f.Close()
 
    s, err := f.RenderUTF8Solid(msg, sdl.Color{
        R: 0xff, G: 0xA0, B: 0xf1, A: 255,
    })
    
    if err != nil{
        return fmt.Errorf("could not render title %v", err)
    }
    defer s.Free()
    t, err :=r.CreateTextureFromSurface(s)
    if err != nil{
        return fmt.Errorf("could not create texture  %v", err)
    }
    defer t.Destroy()
    if err = r.Copy(t, nil, &sdl.Rect{X:(SCREEN_WIDTH /2)/2, Y:(SCREEN_HIGHT /2)/2, W: SCREEN_WIDTH /2, H:SCREEN_HIGHT /6 }); err != nil{
        return fmt.Errorf("could not copy texture  %v", err)
    }
    r.Present()
     
    
    return nil
}
 