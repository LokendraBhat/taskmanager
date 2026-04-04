provider "vagrant" {
 
}

resource "vagrant_vm" "my_vagrant_vm" {
  env = {
    # force terraform to re-run vagrant if the Vagrantfile changes
    VAGRANTFILE_HASH = md5(file(var.vfilename)),
  }
  get_ports = true
  id = 101
  vagrantfile_dir=/home/lokendra/vagrant
  # other
  # name="abc"
  # env=
}
