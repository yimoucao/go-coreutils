mkdir -v -p abc // normal flag
mkdir -vp abc // combined flag
mkdir -vpm abc // combined flag // TOCHECK
mkdir -m mode -vp abc // TO CHECK

mkdir -x abc // other flag
mkdir --p abc // double dash
mkdir -m abc // no mode specified; or invalid mode
mkdir abc abc/1 // one by one creating
mkdir -v a b c // show all verbose
mkdir abc1 abc2 abc3; mkdir abc1 abc2 abc3 // show all existed