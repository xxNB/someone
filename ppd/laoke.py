import pandas as pd 

df = pd.read_excel('/Users/zhangxin/Desktop/ppdai/2drools/2drools/20191017_老客block逾期系数.xlsx')



class Data:
    df=df

# print(df)
# 初始化columns
def reanme_df():
    dic={}
    for i in df.columns:
        print(i)
        dic[i] = str(int(i.split(':')[1].strip(' '))-2)
    Data.df=df.rename(columns=dic)
    # print(Data.df)

def write_mysql(df):
     for x in df.index:
        for y in df.columns:
            print('===========', x+1, y, df.loc[x,y])

def all_due():
    df = Data.df.loc[1:11, ['0','1','2', '3',  '4' , '5','6','7', '8','9','10']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)



def due_3():
    df = Data.df.loc[15:25, ['0','1','2', '3',  '4' , '5','6','7', '8','9','10']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

def due_6():   
    df = Data.df.loc[29:39, ['0','1','2', '3',  '4' , '5','6','7', '8','9','10']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

def due_9():   
    df = Data.df.loc[43:53, ['0','1','2', '3',  '4' , '5','6','7', '8','9','10']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

def due_12():   
    df = Data.df.loc[58:68, ['0','1','2', '3',  '4' , '5','6','7', '8','9','10']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

if __name__ == "__main__":
    reanme_df()
    # all_due()
    due_12()
    # reanme_df()