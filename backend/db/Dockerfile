FROM mysql:8.0

EXPOSE 3306

# MySQL設定ファイルをイメージ内にコピー
COPY ./my.cnf /etc/mysql/conf.d/my.cnf

# 初期データ投入ファイル
COPY ./initdb.d/initial.sql /docker-entrypoint-initdb.d/

# 実行
CMD ["mysqld"]