echo "building rust..."
cd rust
cargo build --release --verbose 
cd ..

echo "building golang"
cd golang
go build -o out/golang -v -x
cd ..

echo "building kotlin"
kotlinc kotlin/Main.kt kotlin/Agent.kt -include-runtime -d kotlin/out/kotlin.jar -verbose