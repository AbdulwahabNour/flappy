package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct{
     bg *sdl.Texture
     bird *bird
     pipes *pipes

}

func newScene(r *sdl.Renderer)(*scene, error){
    t, err := img.LoadTexture(r, BACKGROUND_DIR)
    if err != nil{
        return nil, fmt.Errorf("could not create texture %v", err)
    }
  
    bird, err := newBird(r)
    if err != nil{
        return nil, err
    }
    pipes, err := newPipes(r)
    if err != nil{
        return nil, err
    }

    return &scene{bg: t, bird:bird, pipes: pipes}, nil
}
func(s *scene) paint(r *sdl.Renderer) error{

    r.Clear()
 
    if  err := r.Copy(s.bg, nil, nil); err != nil{
       return fmt.Errorf("could not copy texture %v", err)
    }
  
    if err:=s.bird.paint(r); err != nil{
        return fmt.Errorf("could not paint birds %v", err)
    }
    if err:=s.pipes.paint(r); err != nil{
        return fmt.Errorf("could not paint birds %v", err)
    }
    r.Present()
    return nil
}

func (s *scene) runBird(r *sdl.Renderer, eChan chan sdl.Event) <-chan error{
  errChan := make(chan error)
  
     go func(){
         defer close(errChan)
        t := time.Tick(30 * time.Millisecond)
        for  {
            select {
            case e := <-eChan:
                 if r := s.handleEvent(e); r{
                       return
                 }
            case <-t:
              
                s.update()
                if s.bird.isDead(){
                    drawTitle(r, "Game Over")
                    time.Sleep(time.Second)
                    s.restart()
                   
                }
                if err := s.paint(r); err != nil{
                    errChan <- err
                }
            }
           
        }
     }()
 
 return errChan
}
func(s *scene)handleEvent(e sdl.Event)bool{
     switch t := e.(type){
       case *sdl.QuitEvent:
        return true
       case *sdl.MouseButtonEvent:
         
       case *sdl.TextInputEvent:
       
            if  t.Text[0]== 32{
                s.bird.jump()
            }
       default:
     
     }
     return false
}
func (s *scene)restart(){
   s.bird.restart()
   s.pipes.restart()
}
func(s *scene) update(){
     s.bird.update()
     s.pipes.update()
     s.pipes.touch(s.bird)
}
func (s *scene) destroy(){
    s.bg.Destroy()
    s.bird.destroy()
    s.pipes.destroy()
    
}