default countHover = 0
label start:
    show bg peach
    "You are in \"label start\""
    call screen zeil_learnings_logo
    "continue"
label end:
    "You are in \"label end\", having hovered [countHover] times"
    return


screen zeil_learnings_logo:
    default hovered = False
    python:
        def increaseHover(n):
            global countHover
            countHover += n
    vbox: 
        align 0.5, 0
        imagebutton:
            action [Return()]
            auto "zeil_%s.png"
            hovered [SetScreenVariable("hovered", True), Function(increaseHover, 1)]
            unhovered [SetScreenVariable("hovered", False)]
        showif hovered:
            frame:
                align 0.5, 0
                offset 0, 10
                text "You are hovering, [countHover]"
