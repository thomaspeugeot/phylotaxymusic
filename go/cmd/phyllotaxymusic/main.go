package main

import (
	"flag"
	"log"
	"strconv"

	phyllotaxymusic_models "github.com/thomaspeugeot/phyllotaxymusic/go/models"
	phyllotaxymusic_stack "github.com/thomaspeugeot/phyllotaxymusic/go/stack"
	phyllotaxymusic_static "github.com/thomaspeugeot/phyllotaxymusic/go/static"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	gongsvg_stack "github.com/fullstack-lang/gongsvg/go/stack"

	gongtone_models "github.com/fullstack-lang/gongtone/go/models"
	gongtone_stack "github.com/fullstack-lang/gongtone/go/stack"

	gongtree_stack "github.com/fullstack-lang/gongtree/go/stack"

	substackcursor_models "github.com/thomaspeugeot/phyllotaxymusic/substackcursor/go/models"
	substackcursor_stack "github.com/thomaspeugeot/phyllotaxymusic/substackcursor/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("phyllotaxymusic: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := phyllotaxymusic_static.ServeStaticFiles(*logGINFlag)

	// setup phyllotaxymusicStack
	phyllotaxymusicStack := phyllotaxymusic_stack.NewStack(r,
		phyllotaxymusic_models.Phylotaxy.ToString(), *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	phyllotaxymusicStack.Probe.Refresh()
	phyllotaxymusicStack.Stage.Checkout()

	gongsvg_stack := gongsvg_stack.NewStack(r, phyllotaxymusic_models.GongsvgStackName.ToString(), "", "", "", true, true)
	gongtree_stack := gongtree_stack.NewStack(r, phyllotaxymusic_models.SidebarTree.ToString(), "", "", "", true, true)
	gongtone_stack := gongtone_stack.NewStack(r, phyllotaxymusic_models.GongtoneStackName.ToString(), "", "", "", true, true)
	cursorStack := substackcursor_stack.NewStack(r, substackcursor_models.Substackcursor.ToString(), "", "", "", false, false)

	// get the only diagram
	parameters := phyllotaxymusic_models.GetGongstructInstancesMap[phyllotaxymusic_models.Parameter](phyllotaxymusicStack.Stage)

	if len(*parameters) == 0 {
		log.Println("")
		// log.Fatalln("")
	}

	f4 := new(gongtone_models.Freqency).Stage(gongtone_stack.Stage)
	f4.Name = "F4"

	notef4 := new(gongtone_models.Note).Stage(gongtone_stack.Stage)
	notef4.Frequencies = append(notef4.Frequencies, f4)
	notef4.Start = 0
	notef4.Duration = 1
	notef4.Velocity = 1

	gongtone_stack.Stage.Commit()

	parameter := (*parameters)["Reference"]

	tree := new(phyllotaxymusic_models.Tree)
	tree.TreeStack = gongtree_stack
	tree.Stage = phyllotaxymusicStack.Stage

	parameterImpl := new(ParameterImpl)
	parameterImpl.parameter = parameter
	parameterImpl.gongsvgStage = gongsvg_stack.Stage
	parameterImpl.gongtoneStage = gongtone_stack.Stage
	parameterImpl.phyllotaxymusicStage = phyllotaxymusicStack.Stage
	parameterImpl.tree = tree
	parameterImpl.substackcursorStage = cursorStack.Stage

	parameter.Impl = parameterImpl
	phyllotaxymusic_models.GeneratorSingloton.Impl = parameterImpl

	cursor := new(substackcursor_models.Cursor).Stage(cursorStack.Stage)
	_ = cursor
	cursorStack.Stage.Commit()

	// connect parameter to cursor for start playing notification
	notifyCh := make(chan bool)
	cursor.SetNotifyChannel(notifyCh)
	parameter.SetNotifyChannel(notifyCh)
	parameter.SetCursor(cursor)

	// wait loop in cursor. Will commit once it receive a notification.
	cursor.WaitForPlayNotifications(cursorStack.Stage)

	// generate other stacks
	parameterImpl.Generate()
	tree.Generate(parameter)

	cursorStack.Stage.Commit()

	log.Printf("%s", "Server ready serve on localhost:"+strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}

}

type ParameterImpl struct {
	gongsvgStage         *gongsvg_models.StageStruct
	gongtoneStage        *gongtone_models.StageStruct
	phyllotaxymusicStage *phyllotaxymusic_models.StageStruct
	parameter            *phyllotaxymusic_models.Parameter
	tree                 *phyllotaxymusic_models.Tree
	substackcursorStage  *substackcursor_models.StageStruct
}

// Generate implements models.GeneratorInterface.
func (parameterImpl *ParameterImpl) Generate() {
	p := parameterImpl.parameter

	p.ComputeShapes(parameterImpl.phyllotaxymusicStage)
	p.GenerateSvg(parameterImpl.gongsvgStage)
	p.GenerateNotes(parameterImpl.gongtoneStage, parameterImpl.gongsvgStage, parameterImpl.phyllotaxymusicStage)
	parameterImpl.tree.Generate(p)
	parameterImpl.phyllotaxymusicStage.Commit()
	parameterImpl.substackcursorStage.Commit()
}

func (parameterImpl *ParameterImpl) OnUpdated(updatedParameter *phyllotaxymusic_models.Parameter) {

	log.Println("OnUpdated", parameterImpl.parameter.InsideAngle, parameterImpl.parameter.SideLength)
	// phyllotaxymusic_svg.GenerateSvg(parameterImpl.gongsvgStage, parameterImpl.phyllotaxymusicStage)

	updatedParameter.ComputeShapes(parameterImpl.phyllotaxymusicStage)
	updatedParameter.GenerateSvg(parameterImpl.gongsvgStage)
	parameterImpl.tree.Generate(updatedParameter)
	updatedParameter.GenerateNotes(parameterImpl.gongtoneStage, parameterImpl.gongsvgStage, parameterImpl.phyllotaxymusicStage)
	parameterImpl.substackcursorStage.Commit()
}