import datetime
import sqlite3
import requests

def updateNotes():
    headers = """
Host: i.weread.qq.com
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36
cookie: ptui_loginuin=1457404238; pt_sms_phone=138******50; RK=Lt/FVcr1UZ; ptcz=49e28c5f2fcd878cdddb78fa35a44c4f69be76d807d9483903c636ca773e5a20; wr_gid=291662781; wr_theme=white; tvfe_boss_uuid=e4d2882e2e9d2a54; pgv_pvid=6888801424; eas_sid=d136H4G4Y4v6r3o0b0R0f78667; fqm_pvqid=c8e4be79-0b1b-4456-8b43-75d689d4fcf6; same_pc=%7B%7D; pac_uid=0_d1b3c91b90967; wr_vid=6033941; wr_pf=0; wr_rt=web%40R1LtRFZ0XAAXe3sZo0u_WL; wr_localvid=d7a32a8065c1215d7a8ab34; wr_name=%E6%BB%95%E4%BD%91%E6%A0%87; wr_avatar=https%3A%2F%2Fthirdwx.qlogo.cn%2Fmmopen%2Fvi_32%2FOeuVu4PwvfqGrvoiaUcGLq327PVPwpcm0qzsJNG5OIJqVpxevRMft1ic9awrUxgl0m7O5EuW1XWLibRpOYSSp4PjQ%2F132; wr_gender=1; wr_skey=5WhDvtfk
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
"""
    headers = dict(x.split(": ", 1) for x in headers.splitlines() if x) # 键值对化

    url = 'https://i.weread.qq.com/book/bookmarklist'

    # 承载书名和划线的容器
    bookList = [] # [tuples (bookId,title,author,cover)]
    notes = [] # [tuples (bookId,marktext)]

    r = requests.get(url=url,headers=headers,verify=False)
    if r.ok:
        data = r.json()
    else:
        raise Exception(r.text)
    for book in data["books"]:
        bookList.append((book["bookId"],book["title"],book["author"],book["cover"]))
    for note in sorted(data["updated"],key=lambda x:x["bookId"]):
        notes.append((note["bookId"],note["markText"],note["chapterUid"],note["createTime"]))

    # 本地化存储 todo：增量更新，不要每次完整插入一遍
    conn = sqlite3.connect('noteData.db')
    c = conn.cursor()

    c.execute('''
    CREATE TABLE IF NOT EXISTS BOOKS
    (
    ID INT PRIMARY KEY     NOT NULL,
    TITLE          TEXT    NOT NULL,
    AUTHOR         VARCHAR(64) NOT NULL,
    COVER          TEXT    NOT NULL
    );
    ''')
    c.execute('''
    CREATE TABLE IF NOT EXISTS NOTES
    (
    BOOKID           INT              NOT NULL,
    MARKTEXT         VARCHAR(255)     NOT NULL,
    chapterUid       INT              NOT NULL,
    createTime       TIMESTAMP,
    FOREIGN KEY (BOOKID) REFERENCES BOOKS(ID)
    );
    ''')

    for book in bookList:
        c.execute(f"INSERT INTO BOOKS VALUES {book}")
    for note in notes:
        c.execute(f"INSERT INTO NOTES VALUES {note}")

    conn.commit()
    conn.close()

def exportMD():
    conn = sqlite3.connect('noteData.db')
    c1 = conn.cursor()
    c2 = conn.cursor()
    c1.execute('SELECT * FROM BOOKS')
    for book in c1:
        c2.execute(f'SELECT MARKTEXT,chapterUid,createTime FROM NOTES WHERE bookID={book[0]}')
        with open(f'notes/{book[1]}.md','w+',encoding='utf8') as f:
            for note in c2:
                f.write(f'# {note[1]}\n**{datetime.datetime.fromtimestamp(note[2])}** \n\n*{note[0]}*\n\n')
            

if __name__=='__main__':
    # updateNotes()
    exportMD()