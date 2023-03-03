default countHover = 0
label start:
    show bg peach
    "You are in \"label start\""
    call screen zeil_learnings_logo
label end:
    "You are in \"label end\", having hovered [countHover] times"
    return


screen zeil_learnings_logo:
    python:
        def increaseHover(n):
            global countHover
            countHover += n
            
    imagebutton:
        align 0.5,0
        action [Hide("dialogTextScreen"), Jump("end")]
        auto "zeil_%s.png"
        hovered [Function(increaseHover,1), Show("dialogTextScreen", dialogText = "You are hovering")]
        unhovered Hide("dialogTextScreen")

screen dialogTextScreen:
    default dialogText = ""
    vbox:
        align 0.5, 0.75
        frame:
            text f"* {dialogText} *"
