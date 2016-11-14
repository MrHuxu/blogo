const gulp = require('gulp'),
      babel = require('gulp-babel'),
      uglify = require('gulp-uglify'),
      minifyCss = require('gulp-minify-css'),
      concat = require('gulp-concat'),
      concatCss = require('gulp-concat-css'),
      rename = require('gulp-rename');

const jsFiles = [
  './assets/javascripts/util.js',
  './assets/javascripts/layout.js',
  './assets/javascripts/*.js'
];

gulp.task('js', function () {
  gulp.src(jsFiles)
    .pipe(concat('all.min.js'))
    .pipe(gulp.dest('./assets/dist'))
    .pipe(babel({ presets: ['es2015'] }))
    .pipe(uglify())
    .pipe(gulp.dest('./assets/dist'));
});

gulp.task('css', function () {
  gulp.src([
      './assets/stylesheets/layout.css',
      './assets/stylesheets/*.css'
    ])
    .pipe(concatCss('all.min.css'))
    .pipe(minifyCss({compatibility: 'ie8'}))
    .pipe(gulp.dest('./assets/dist'));
});

gulp.task('compress', ['js', 'css']);
