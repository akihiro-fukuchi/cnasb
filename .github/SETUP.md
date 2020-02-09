
gen ssh key
```
ssh-keygen -t rsa -C "<email>" -f github.com.id_rsa
```

use my github.com id_rsa
```
cat << EOS >> ~/.ssh/config
Host github.com
  User git
  HostName github.com
  IdentityFile ~/.ssh/github.com.id_rsa
  IdentitiesOnly yes
EOS
```

set gitconfig
```
git config --local user.name '<name>'
git config --local user.email '<email>'
```

check condition
```
ssh -T git@github.com
```

