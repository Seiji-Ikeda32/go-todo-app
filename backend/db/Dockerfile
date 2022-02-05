FROM mysql:8.0

EXPOSE 3306

# MySQL設定ファイルをイメージ内にコピー
ADD ./my.cnf /etc/mysql/conf.d/my.cnf

# docker runで実行
# CMD ["mysqld"]