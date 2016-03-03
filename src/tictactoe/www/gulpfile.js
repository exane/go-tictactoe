var browserify = require('browserify');
var gulp = require('gulp');
var source = require('vinyl-source-stream');
var babelify = require("babelify");
var watch = require("gulp-watch");

//fast install
//npm i --save-dev browserify gulp vinyl-source-stream babelify


gulp.task('browserify', function(){
  browserify('./js/main.js', {standalone: "app", debug: true})
  .transform("babelify", {
    presets: ["es2015", "stage-0"]
  })
  .bundle().on("error", function(err){
    console.log(err);
  })
  .pipe(source('app.js').on("error", function(err){
    console.log(err);
  }))
  .pipe(gulp.dest('./build/').on("error", function(err){
    console.log(err);
  }));

});


gulp.task("watch", function(){
  watch("./js/**/*.js", function(){
    gulp.start("browserify");
  });
})


gulp.task("default", ["watch", "browserify"]);