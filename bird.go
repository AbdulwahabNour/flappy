package main

import (
	"fmt"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)
const(
    BIRD_FRAME_DIR = "res/imgs/bird_frame_"
    BIRD_FRAME_EXT = ".png"
    GRAVITY = 0.25
)

type bird struct{
        mu sync.RWMutex
        textures []*sdl.Texture
        time int
        x,y int32
        w, h int32
        speed float64
        dead bool

}

func newBird(r *sdl.Renderer) (*bird, error){
   t := make([]*sdl.Texture,4)
    for i := 0; i < 4; i++ {
         b , err := img.LoadTexture(r, fmt.Sprint(BIRD_FRAME_DIR, (i+1), BIRD_FRAME_EXT))
         if err != nil{
             return nil, fmt.Errorf("could not create texture %v", err)
         }
         t[i] = b
    }

    return &bird{textures: t,x: 10, y: 400,w: 50, h: 43}, nil
}

func(b *bird)paint(r *sdl.Renderer) error{
    b.mu.RLock() 
    defer b.mu.RUnlock()

    b.time++
    rec := &sdl.Rect{X: b.x, Y:  ((SCREEN_HIGHT -  b.y  )  - b.h /2 )  , W: b.w, H: b.h}
    i := b.time 
    if  err := r.Copy(b.textures[i], nil, rec); err != nil{
        return fmt.Errorf("could not copy texture %v", err)
     }

     if b.time >= 3{
        b.time = 0
    }

 
 return nil
}
func(b *bird)update(){
    b.mu.Lock()
    defer b.mu.Unlock()
    b.y -= int32(b.speed)
    if b.y < 0 {
           b.dead = true
    }
    b.speed += GRAVITY
  
  
} 
func (b *bird)restart(){
    b.mu.Lock()
    defer b.mu.Unlock()
    b.y = 400
    b.speed = 0
    b.dead = false
  
}
func (b *bird) destroy(){
     b.mu.Lock()
     defer b.mu.Unlock()
     for _, v:= range b.textures{
          v.Destroy()
     }
    
}
func(b *bird)isDead()bool{
    return b.dead
}

func(b *bird)jump(){
    b.mu.Lock()
    defer b.mu.Unlock()
    b.speed = -5
   
}
func(b *bird) touch(p *pipe){
    b.mu.Lock()
    defer b.mu.Unlock()
    if p.x > b.x+b.w{ 
        return
    }
    if p.x + p.w < p.x{
        return
    }
    if !p.p && p.h < b.y - b.h/2{
        return
    }
    if p.p && SCREEN_HIGHT- p.h > b.y + b.h/2{
        return
    }
    b.dead = true
}