const gulp = require('gulp'),
  babel = require('gulp-babel'),
  uglify = require('gulp-uglify'),
  cleanCss = require('gulp-clean-css'),
  concat = require('gulp-concat'),
  concatCss = require('gulp-concat-css'),
  rename = require('gulp-rename');

const jsFiles = [
  './app/assets/javascripts/util.js',
  './app/assets/javascripts/layout.js',
  './app/assets/javascripts/*.js'
];

gulp.task('js', function () {
  gulp.src(jsFiles)
    .pipe(concat('all.min.js'))
    .pipe(gulp.dest('./app/assets/dist'))
    .pipe(babel({ presets: ['env'] }))
    .pipe(uglify())
    .pipe(gulp.dest('./app/assets/dist'));
});

gulp.task('css', function () {
  gulp.src([
    './app/assets/stylesheets/layout.css',
    './app/assets/stylesheets/*.css'
  ])
    .pipe(concatCss('all.min.css'))
    .pipe(cleanCss({ compatibility: 'ie8' }))
    .pipe(gulp.dest('./app/assets/dist'));
});

gulp.task('compress', ['js', 'css']);
