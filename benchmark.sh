hyperfine --warmup 3 \
    "./rust/target/release/rust" --command-name "Rust" \
    "./golang/out/golang" --command-name "Golang" \
    "java -jar kotlin/out/kotlin.jar" --command-name "Kotlin (JVM)" \
    "python3 python/main.py" --command-name "Pure Python"