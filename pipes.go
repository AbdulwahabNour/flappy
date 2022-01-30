package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)
const(
    PIPE_IMAGE_DIR = "res/imgs/pipe.png"
)
type pipes struct{
    mu sync.RWMutex
    textures *sdl.Texture
    speed int32
    pipes []*pipe
}
type pipe struct{
    mu sync.RWMutex
     x int32
     w int32
     h int32
     p bool

}
func newPipes(r *sdl.Renderer)(*pipes, error){

    p , err := img.LoadTexture(r, PIPE_IMAGE_DIR)
    if err != nil{
            return nil, fmt.Errorf("could not create texture %v", err)
    }
npipes:= &pipes{ speed: 3, textures: p}
   go func(){
          for{
                npipes.mu.Lock()
         
                npipes.pipes = append(npipes.pipes,  newPipe())
                npipes.mu.Unlock()
                time.Sleep(2 *time.Second)
              
          }
   }()
    return npipes , nil
}
func(p *pipes)paint(r *sdl.Renderer)error{
    p.mu.RLock()
    defer p.mu.RUnlock()

    for _, v := range p.pipes{
        if  err := v.paint(r, p.textures);err != nil{
            return err
        }
    }  
  
 return nil
 }
 
 func(p *pipes)restart(){
     p.mu.Lock()
     defer p.mu.Unlock()
     p.pipes = nil
 }
 func(p *pipes)update(){
     p.mu.Lock()
     defer p.mu.Unlock()
     var rem []*pipe
     for _, v := range p.pipes{
         v.mu.Lock()
        v.x -= p.speed
        v.mu.Unlock()
        if v.x + v.w >0 {
             rem = append(rem, v)
        }
     }
     p.pipes = rem
     
 }
 func(p *pipes) touch(b *bird){
    p.mu.RLock()
    defer p.mu.RUnlock()

    for _, v:= range p.pipes{
         b.touch(v)
    }
 
}
 func(p *pipes)destroy(){
     p.mu.Lock()
     defer p.mu.Unlock()
     p.textures.Destroy()
 }
 

func newPipe()*pipe{
         return &pipe{   x:1000,
                        w: 50,
                        h: 150 + int32(rand.Intn(300)),
                        p: rand.Float32() > 0.5,
                           }
}


func(p *pipe)paint(r *sdl.Renderer, texture *sdl.Texture)error{
   p.mu.RLock()
   defer p.mu.RUnlock()
    rec := &sdl.Rect{X: p.x, Y: SCREEN_HIGHT -p.h , W: p.w, H:p.h}
    f := sdl.FLIP_NONE
    if p.p {
         rec.Y =0
          f = sdl.FLIP_VERTICAL
    }
   
    if  err :=  r.CopyEx(texture, nil, rec, 0, nil, f); err != nil{
        return fmt.Errorf("could not copy texture %v", err)
     }

 

return nil
}
func(p *pipe)toch(b *bird){
    p.mu.RLock()
    defer p.mu.RUnlock()
    b.touch(p)
}


 
 
 
