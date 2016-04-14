node {

   // checkout
   stage 'Checkout'
   git url: 'https://github.com/tommymalmqvist/golang-hourglass.git'

   // build
   stage 'Build'
   sh "make build"

   // test
   stage 'Test'
   sh "make test"

   // build
   stage 'Clean'
   sh "make clean"
}
