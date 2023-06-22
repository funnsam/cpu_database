module github.com/funnsam/cpu_db/console

go 1.19

replace github.com/funnsam/cpu_db => ../
replace github.com/funnsam/cpu_db/reader => ../reader

require github.com/funnsam/cpu_db/reader/src v0.0.0-00010101000000-000000000000 // indirect
