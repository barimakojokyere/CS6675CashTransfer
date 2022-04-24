echo 'Building client binary...'
go build -o bin/client srvcs/client/clntmain.go
sleep 2
echo 'Done'
echo 'Building Server binary...'
go build -o bin/server srvcs/server/svrmain.go
sleep 2
echo 'Done'
echo 'Building MoMo server binary...'
go build -o bin/momoserver srvcs/momoserver/momosvrmain.go
sleep 2
echo 'Done'
echo 'Building PayPal server binary...'
go build -o bin/paypalserver srvcs/paypalserver/paypalsvrmain.go