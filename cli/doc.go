package cli

/*
export SEVENCONFIG=./config/sevenconfig.yaml 

# prompt with human question only

go run main.go apply \
--manifest robot/01-simple.yaml \
--envfile ollama.env \
--logs \
--output output.txt

go run main.go apply \
--manifest robot/01-simple.yaml \
--envfile ollama.env \
--logs \
--output output.txt \
--question "[brief] who is Jean-Luc Picard"

go run main.go apply \
--manifest robot/01-simple.yaml \
--envfile ollama.env \
--logs \
--output output.txt \
--question "Who is Michael Burnham?"


# prompt with human question and system instruction
go run main.go apply \
--manifest robot/02-prompt-system.yaml \
--envfile ollama.env \
--logs \
--output output.txt

go run main.go apply \
--manifest robot/02-prompt-system.yaml \
--envfile ollama.env \
--logs \
--output output.txt \
--system "override the system message"


go run main.go apply \
--manifest robot/02-prompt-system.yaml \
--envfile ollama.env \
--logs \
--output output.txt \
--question "Who is Michael Burnham?"


# prompt with human question, system instruction and context

go run main.go apply \
--manifest robot/03-prompt-system-context.yaml \
--envfile ollama.env \
--logs \
--output output.txt

go run main.go apply \
--manifest robot/03-prompt-system-context.yaml \
--envfile ollama.env \
--logs \
--question "Who is Jean-Luc Picard?"

go run main.go apply \
--manifest robot/03-prompt-system-context.yaml \
--envfile ollama.env \
--logs \
--question "What is your name?"

go run main.go apply \
--manifest robot/03-prompt-system-context.yaml \
--envfile ollama.env \
--logs \
--question "Who is the best friend of James Kirk?"

go run main.go apply \
--manifest robot/03-prompt-system-context.yaml \
--envfile ollama.env \
--logs \
--output output.txt \
--context "override the context message"

go run main.go apply \
--manifest robot/03-prompt-system-context.yaml \
--envfile ollama.env \
--logs \
--output output.txt \
--context "override the context message" \
--system "override the system message"

# prompt with documents

go run main.go apply \
--manifest robot/04-prompt-documents.yaml \
--envfile ollama.env \
--logs
*/