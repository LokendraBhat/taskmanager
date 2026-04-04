provider "vagrant" {
 
}

resource "vagrant_vm" "vagrant_vm_alma" {
  env = {
    VAGRANTFILE_HASH = md5(file(var.vfile_alma))
  }
  get_ports       = true
  vagrantfile_dir = "/home/lokendra/vagrant/alma"
}

resource "vagrant_vm" "vagrant_vm_debian" {
  env = {
    VAGRANTFILE_HASH = md5(file(var.vfile_debian))
  }
  get_ports       = true
  vagrantfile_dir = "/home/lokendra/vagrant/debian"
}