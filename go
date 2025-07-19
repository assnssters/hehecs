sudo ls /dev 
sudo lsblk
read -p "choose disk /dev/" disk
cd ~/
mkdir diska
sudo mount /dev/$disk diska && echo "OK:Disk is mounted!" || echo "E:Mount disk error 404/1"
sudo sleep 2
cd ~/diska 
ls
read -p "Sure this disk is filesystem host? y/n: " yan
case $yan in
    Y|y) clear; echo "Continue to go host!";;
    N|n) echo "Pls try again, huhu";exit;;
    *) echo "bro wtf huhu";exit;;
esac
ssh-keygen -t rsa -b 2048 -N "" -f ~/ssh
sudo cat ~/ssh.pub >> ~/diska/root/.ssh/authorized_keys
clear
sleep 4
ssh -i ~/ssh -o StrictHostKeyChecking=accept-new root@localhost -p 2000
  ssh-keygen -f '/root/.ssh/known_hosts' -R '[localhost]:2000'
chmod +x log
