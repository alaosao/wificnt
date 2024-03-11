<div align="center">

# NvArch

</div>

<p align="center">
    <a href="#">
<img alt="Made with Shell" src="https://img.shields.io/badge/Shell_Script-121011?style=for-the-badge&logo=gnu-bash&logoColor=white">
    </a>
</p>

Written purely in shell, networkmanager program made to automate and provide seamless connection to known networks.

<details>
    <summary><strong>Guide</strong></summary>

Generally, your SSID and password does not save, leading you to constantly retype your Network's password, creating a tiring process. This shell project is meant to change that. 

**wifirand** <SSID>:
Typing the wifirand command will lead to the program to save your SSID and password. However, please keep in mind that wifirand must be accompanied by your <SSID>. It will prompt you to save a password. 

ex: `wifirand Verizon_XXXX`
*Keep in mind that as of writing this, only one SSID and password can be saved at a time*

**wificon**:
Typing this command will auto connect you to the SSID and password you listed in the `wifirand` command.

ex: `wificon`

**Where does it save?**:

It saves to a hidden file in your home directory, specifically, `~/.wificnt/wifi.md`. Through here, the shell script is able to edit and save your wifi in plaintext. 

</details>

# Installation
**[NETWORKMANAGER](https://archlinux.org/packages/extra/x86_64/networkmanager/) MUST BE INSTALLED**

## Option 1: Shell

- Install the repository

```bash
git clone https://github.com/mitzsou/wificnt.git
```

- Enter the directory

```bash
cd ~/wificnt/
```

- Elevate the script's permissions and run it!

```bash
chmod +x setup.sh 
./setup.sh
```

## Option 2: Manual

- Install the dependencies

### Arch Linux
```bash
sudo pacman -S networkmanager git
```
### Debian/Ubuntu
```bash
sudo apt install network-manager git
```

### openSUSE
```bash
sudo zypper install NetworkManager git
```

### Gentoo
```bash
sudo emerge net-misc/networkmanager dev-vcs/git
```

- Setup NetworkManager
```bash
sudo systemctl start NetworkManager
sudo systemctl enable NetworkManager
```

- Install the repository

```bash
git clone https://github.com/mitzsou/wificnt.git
```

- Create the directory & save file

```bash
mkdir ~/.wificnt/
touch ~/.wificnt/wifi.md 
```

- Move the designated files to your system's path

```bash
mv ~/wificnt/wificon /usr/local/bin
mv ~/wificnt/wificon /usr/local/bin
``` 
*Your path may be different*

# To-do 

- Include a way to save multiple SSID's and SSID password's at once. 
- Include functionality for all linux distrobutions for setup.sh
- think of more stuff to do w this lol
