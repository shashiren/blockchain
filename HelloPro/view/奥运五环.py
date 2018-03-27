import turtle

turtle.color("black")
turtle.circle(100)
#turtle.done()

turtle.penup()#抬起笔
turtle.goto(-200,0)#移动到位置
turtle.pendown()#笔放下
turtle.color("blue")#颜色
turtle.circle(100)#半径100

turtle.penup()
turtle.goto(200,0)
turtle.pendown()
turtle.color("red")
turtle.circle(100)

turtle.penup()
turtle.goto(100,-100)
turtle.pendown()
turtle.color("green")
turtle.circle(100)

turtle.penup()
turtle.goto(-100,-100)
turtle.pendown()
turtle.color("yellow")
turtle.circle(100)
turtle.done()#程序继续执行