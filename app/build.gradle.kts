plugins {
    kotlin("jvm") version "2.0.10"
    application
    antlr
}

repositories {
    mavenCentral()
}

dependencies {
    testImplementation(libs.junit.jupiter)
    testRuntimeOnly("org.junit.platform:junit-platform-launcher")
    implementation(libs.guava)
    implementation("org.jsoup:jsoup:1.18.1")
    antlr("org.antlr:antlr4:4.13.2")

}

java {
    toolchain {
        languageVersion = JavaLanguageVersion.of(21)
    }
}

application {
    mainClass = "io.htql.Htql"
}

tasks.named<Test>("test") {
    useJUnitPlatform()
}

tasks {

    generateGrammarSource {
        arguments = arguments + listOf(
            "-visitor",
            "-package", "io.htql.parser"
        )
        outputDirectory = file("src/main/java/io/htql/parser")
    }

    compileKotlin {
        dependsOn(generateGrammarSource)
    }

    compileTestKotlin {
        dependsOn(generateTestGrammarSource)
    }
    jar {
        manifest {
            attributes["Main-Class"] = application.mainClass
        }
    }
}