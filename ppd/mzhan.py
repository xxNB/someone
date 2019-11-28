import pandas as pd 

df = pd.read_excel('/Users/zhangxin/Desktop/ppdai/2drools/2drools/新客M站逾期映射表.xlsx')



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
    df = Data.df.loc[3:22, ['1','2', '3',  '4' , '5']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)



def due_3():
    df = Data.df.loc[30:49, ['1','2', '3',  '4' , '5']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

def due_6():   
    df = Data.df.loc[53:72, ['1','2', '3',  '4' , '5']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

def due_9():   
    df = Data.df.loc[77:96, ['1','2', '3',  '4' , '5']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

def due_12():   
    df = Data.df.loc[102:121, ['1','2', '3',  '4' , '5']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

if __name__ == "__main__":
    reanme_df()
    all_due()
    # due_9()
    # reanme_df()